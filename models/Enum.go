package models

import (
	"github.com/optim-corp/cios-cli/utils"
)

var (
	LIST         = "list"
	PATCH        = "patch"
	CREATE       = "create"
	DELETE       = "delete"
	ALIAS_LIST   = []string{"ls", "show", "out", "output"}
	ALIAS_PATCH  = []string{"patch", "update", "up", "pct", "upgrade", "upd"}
	ALIAS_CREATE = []string{"cre", "add", "sub", "subscribe"}
	ALIAS_DELETE = []string{"del", "rm", "remove", "clean", "clear", "cle"}

	Stages    = createStages()
	FullScope = "openid email profile address user.profile user.read user.write " +
		"group.read group.write group.relation.read group.relation.write" +
		"corporation.read corporation.write corporation.user.read corporation.user.write corporation.group.read  corporation.group.write " +
		"oauth2_client.read oauth2_client.write license.read license.write acl.read acl.write resource_owner.read " +
		"resource_owner.read " +
		"channel_protocol.read channel_protocol.write.internal channel.read  channel.write  " +
		"messaging.subscribe  messaging.publish " +
		"datastore.read  datastore.write  datastore.upload  datastore.download " +
		"file_storage.read file_storage.write file_storage.upload  file_storage.download  " +
		"device.read  device.write  device_model.read device_model.write device.group_policy.read devices.group_policies.write " +
		"geo.area.read geo.area.write geo.area-kind.read geo.area-kind.write geo.area.content.write geo.circle.read geo.circle.write geo.map.read geo.map.write geo.point.read geo.point.write geo.polygon.read geo.polygon.write geo.route.read geo.route.write " +
		"active_license.read active_license.write product.read " +
		"videostream.read videostream.view "
	URL_JSON = `{
			"alstroemeria": {
				"DeviceManagement": "device-management.optimcloudapis.com",
				"DeviceAssetManagement": "device-asset-lifecycle.optimcloudapis.com",
				"Monitoring": "monitoring.optimcloudapis.com",
				"Messaging": "messaging.optimcloudapis.com",
				"Location": "location.optimcloudapis.com",
				"Accounts": "accounts.optimcloudapis.com",
				"Storage": "storage.optimcloudapis.com",
				"Iam": "iam.optimcloudapis.com",
				"Auth": "auth.optim.cloud",
				"VideoStreams": "video-streams.optim.cloud"
			},
			"viola": {
                "DeviceManagement": "device-management.preapis.cios.dev",
                "Monitoring": "monitoring.preapis.cios.dev",
                "Messaging": "messaging.preapis.cios.dev",
                "Location": "location.preapis.cios.dev",
                "Accounts": "accounts.preapis.cios.dev",
                "Storage": "storage.preapis.cios.dev",
                "Iam": "iam.preapis.cios.dev",
                "Auth": "auth.pre.cios.dev",
                "VideoStreams": "video-streaming.preapis.cios.dev"
        	}
		}`
)

func createStages() []string {
	if str, err := utils.Path(utils.UrlPath).ReadString(); err != nil {
		return utils.GetKeys(utils.StrToMap(URL_JSON))
	} else {
		return utils.GetKeys(utils.StrToMap(str))
	}
}
