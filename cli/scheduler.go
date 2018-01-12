package cli

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
	"go.uber.org/zap"

	ds "github.com/sonm-io/marketplace/datastruct"
)

func (a *App) runScheduler() error {
	cleanUpPeriod, err := time.ParseDuration(a.conf.OrdersCleanUpPeriod)
	if err != nil {
		return fmt.Errorf("orders_cleanup_period has invalid value %q", a.conf.OrdersCleanUpPeriod)
	}

	ordersTTL, err := time.ParseDuration(a.conf.OrdersTTL)
	if err != nil {
		return fmt.Errorf("orders_ttl has invalid value %q", a.conf.OrdersTTL)
	}

	a.logger.Info("Starting scheduler")
	a.schedulerQuitCh = gocron.Start()

	a.logger.Info("Expired orders clean up period is " + cleanUpPeriod.String())
	gocron.Every(uint64(cleanUpPeriod.Seconds())).Seconds().Do(a.cleanExpiredAskOrders(ordersTTL))

	return nil
}

func (a *App) stopScheduler() {
	a.logger.Info("Stopping scheduler")
	a.schedulerQuitCh <- true
	a.logger.Info("Scheduler stopped")
}

func (a *App) cleanExpiredAskOrders(TTL time.Duration) func() {
	return func() {
		a.logger.Info("Cleaning orders")
		a.logger.Info("Orders TTL is " + TTL.String())

		// status 3 = EXPIRED ORDER
		sql := `UPDATE orders
			    SET status=?
				WHERE type=?
					AND status != 3
					AND (strftime('%s', 'now') - strftime('%s', updated_at)) > ?`

		res, err := a.db.Exec(sql, 3, ds.Ask, uint64(TTL.Seconds()))
		if err != nil {
			a.logger.Warn("Cannot clean expired orders", zap.Error(err))
		}

		num, _ := res.RowsAffected()
		a.logger.Info(fmt.Sprintf("Cleaned orders: %d", num))
	}
}
