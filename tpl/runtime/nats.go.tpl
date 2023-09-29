/*
 * Copyright (C) ###__PROJ_AUTHOR__### - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file nats.go
 * @package runtime
 * @author ###__PROJ_AUTHOR__###
 * @since ###__TODAY__###
 */

package runtime

import "github.com/nats-io/nats.go"

var Nats *nats.Conn

func InitNats() error {
	nc, err := nats.Connect(Config.Nats.URL)
	if err != nil {
		Logger.Fatal("nats connection failed : %s", err)
	}

	Nats = nc

	return err
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
