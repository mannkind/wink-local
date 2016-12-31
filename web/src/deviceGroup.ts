import Vue = require("vue");
import axios, { Promise } from "axios";
import * as vts from "vue-typescript-component";
import * as AddDevice from "./device/addDevice.vue";
import * as ListDevice from "./device/listDevice.vue";
import * as AddGroup from "./group/addGroup.vue";
import * as ListGroup from "./group/listGroup.vue";
import { ACTIONS } from "./models/actions";
import { Apron } from "./models/apron";
import { IApronDeviceGroup } from "./models/apronDeviceGroup";
import { EventBus } from "./models/events";

@vts.component({ components: { AddDevice, ListDevice, AddGroup, ListGroup } })
export default class DeviceGroup extends Vue {
    public devices: IApronDeviceGroup[] = [];
    public groups: IApronDeviceGroup[] = [];
    private pairingTimeout: number = 60000;

    public created() {
        EventBus.$on(ACTIONS.ADD_DEVICE, this.addDevice);
        EventBus.$on(ACTIONS.REMOVE_DEVICE, this.removeDevice);
        EventBus.$on(ACTIONS.UPDATE_DEVICE, this.updateDevice);
        EventBus.$on(ACTIONS.ADD_GROUP, this.addGroup);
        EventBus.$on(ACTIONS.REMOVE_GROUP, this.removeGroup);
        EventBus.$on(ACTIONS.ADD_DEVICE_TO_GROUP, this.addDeviceToGroup);
        EventBus.$on(ACTIONS.REMOVE_DEVICE_FROM_GROUP, this.removeDeviceFromGroup);

        this.updateDeviceList();
        this.updateGroupList();
    }

    private addDevice(radio: string) {
        Apron.addDevice(radio);

        setTimeout(() => {
            this.updateDeviceList();
            EventBus.$emit(ACTIONS.ADDED_DEVICE);
        }, this.pairingTimeout);
    }

    private removeDevice(deviceId: number) {
        Apron.removeDevice(deviceId).then((_) => {
            this.devices = this.devices.filter((device) => device.ID !== deviceId);
            this.groups.forEach((x) => {
                if (x.Nodes == null) {
                    return;
                }

                const index = x.Nodes.findIndex((y) => y.ID === deviceId);
                if (index === -1) {
                    return;
                }

                x.Nodes.splice(index, 1);
            });
            EventBus.$emit(ACTIONS.REMOVED_DEVICE);
        });
    }

    private updateDevice(deviceId: number, name: string) {
        const done = () => { EventBus.$emit(ACTIONS.UPDATED_DEVICE); };
        const device = this.devices.find((x) => x.ID === deviceId);
        if (device == null) {
            done();
            return;
        }

        Apron.updateDevice(deviceId, name).then((response) => {
            device.Name = name;
            this.groups.forEach((x) => {
                if (x.Nodes == null) {
                    return;
                }

                const groupDevice = x.Nodes.find((y) => y.ID === deviceId);
                if (groupDevice == null) {
                    return;
                }

                groupDevice.Name = name;
            });
            done();
        });
    }

    private addDeviceToGroup(groupId: number, deviceId: number) {
        const done = () => { EventBus.$emit(ACTIONS.ADDED_DEVICE_TO_GROUP); };
        const group = this.groups.find((x) => x.ID === groupId);
        const node = this.devices.find((x) => x.ID === deviceId);

        Apron.addDeviceToGroup(groupId, deviceId).then((response) => {
            if (group == null || group.Nodes == null || node == null) {
                done();
                return;
            }

            group.Nodes.push(node);
            done();
        });
    }

    private removeDeviceFromGroup(groupId: number, deviceId: number) {
        const done = () => { EventBus.$emit(ACTIONS.REMOVED_DEVICE_FROM_GROUP); };
        const group = this.groups.find((x) => x.ID === groupId);

        Apron.removeDeviceFromGroup(groupId, deviceId).then((response) => {
            if (group == null || group.Nodes == null) {
                done();
                return;
            }

            group.Nodes = group.Nodes.filter((node) => node.ID !== deviceId);
            done();
        });
    }

    private addGroup(name: string) {
        Apron.addGroup(name).then((_) => {
            this.updateGroupList().then((__) => {
                EventBus.$emit(ACTIONS.ADDED_GROUP);
            });
        });
    }

    private removeGroup(groupId: number) {
        Apron.removeGroup(groupId).then((_) => {
            this.groups = this.groups.filter((group) => group.ID !== groupId);
            EventBus.$emit(ACTIONS.REMOVED_GROUP);
        });
    }

    private updateDeviceList(): Promise<void> {
        // console.log("Updated devices");
        return axios
            .get("/device/list")
            .then((response) => {
                // console.log("Updated devices");
                this.devices = response.data;
            })
            .catch((reason) => {
                // console.log("Failed to fetch device list");
            });
    }

    private updateGroupList(): Promise<void> {
        // console.log("Updating groups");
        return axios
            .get("/group/list")
            .then((response) => {
                // console.log("Updated groups");
                this.groups = response.data;
            })
            .catch((reason) => {
                // console.log("Failed to fetch device list");
            });
    }
}
