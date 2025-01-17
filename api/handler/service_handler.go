// Copyright (C) 2014-2018 Goodrain Co., Ltd.
// RAINBOND, Application Management Platform

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package handler

import (
	api_model "github.com/goodrain/rainbond/api/model"
	"github.com/goodrain/rainbond/api/util"
	"github.com/goodrain/rainbond/builder/exector"
	dbmodel "github.com/goodrain/rainbond/db/model"
	"github.com/goodrain/rainbond/worker/discover/model"
	"github.com/goodrain/rainbond/worker/server/pb"
)

//ServiceHandler service handler
type ServiceHandler interface {
	ServiceBuild(tenantID, serviceID string, r *api_model.BuildServiceStruct) error
	AddLabel(l *api_model.LabelsStruct, serviceID string) error
	DeleteLabel(l *api_model.LabelsStruct, serviceID string) error
	UpdateLabel(l *api_model.LabelsStruct, serviceID string) error
	UpdateServiceLabel(serviceID, value string) error
	StartStopService(s *api_model.StartStopStruct) error
	ServiceVertical(v *model.VerticalScalingTaskBody) error
	ServiceHorizontal(h *model.HorizontalScalingTaskBody) error
	ServiceUpgrade(r *model.RollingUpgradeTaskBody) error
	ServiceCreate(ts *api_model.ServiceStruct) error
	ServiceUpdate(sc map[string]interface{}) error
	LanguageSet(langS *api_model.LanguageSet) error
	GetService(tenantID string) ([]*dbmodel.TenantServices, error)
	GetPagedTenantRes(offset, len int) ([]*api_model.TenantResource, int, error)
	GetTenantRes(uuid string) (*api_model.TenantResource, error)
	CodeCheck(c *api_model.CheckCodeStruct) error
	ServiceDepend(action string, ds *api_model.DependService) error
	EnvAttr(action string, at *dbmodel.TenantServiceEnvVar) error
	PortVar(action string, tenantID, serviceID string, vp *api_model.ServicePorts, oldPort int) error
	PortOuter(tenantName, serviceID string, containerPort int, servicePort *api_model.ServicePortInnerOrOuter) (*dbmodel.TenantServiceLBMappingPort, string, error)
	PortInner(tenantName, serviceID, operation string, port int) error
	VolumnVar(avs *dbmodel.TenantServiceVolume, tenantID, fileContent, action string) *util.APIHandleError
	UpdVolume(sid string, req *api_model.UpdVolumeReq) error
	VolumeDependency(tsr *dbmodel.TenantServiceMountRelation, action string) *util.APIHandleError
	GetDepVolumes(serviceID string) ([]*dbmodel.TenantServiceMountRelation, *util.APIHandleError)
	GetVolumes(serviceID string) ([]*dbmodel.TenantServiceVolume, *util.APIHandleError)
	ServiceProbe(tsp *dbmodel.TenantServiceProbe, action string) error
	RollBack(rs *api_model.RollbackStruct) error
	GetStatus(serviceID string) (*api_model.StatusList, error)
	GetServicesStatus(tenantID string, services []string) map[string]string
	CreateTenant(*dbmodel.Tenants) error
	CreateTenandIDAndName(eid string) (string, string, error)
	GetPods(serviceID string) (*K8sPodInfos, error)
	TransServieToDelete(serviceID string) error
	TenantServiceDeletePluginRelation(tenantID, serviceID, pluginID string) *util.APIHandleError
	GetTenantServicePluginRelation(serviceID string) ([]*dbmodel.TenantServicePluginRelation, *util.APIHandleError)
	SetTenantServicePluginRelation(tenantID, serviceID string, pss *api_model.PluginSetStruct) (*dbmodel.TenantServicePluginRelation, *util.APIHandleError)
	UpdateTenantServicePluginRelation(serviceID string, pss *api_model.PluginSetStruct) (*dbmodel.TenantServicePluginRelation, *util.APIHandleError)
	UpdateVersionEnv(uve *api_model.SetVersionEnv) *util.APIHandleError
	DeletePluginConfig(serviceID, pluginID string) *util.APIHandleError
	ServiceCheck(*api_model.ServiceCheckStruct) (string, string, *util.APIHandleError)
	GetServiceCheckInfo(uuid string) (*exector.ServiceCheckResult, *util.APIHandleError)
	GetServiceDeployInfo(tenantID, serviceID string) (*pb.DeployInfo, *util.APIHandleError)
	ListVersionInfo(serviceID string) (*api_model.BuildListRespVO, error)
}
