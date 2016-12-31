import Vue = require("vue");
import * as vts from "vue-typescript-component";
import { ACTIONS } from "../models/actions";
import { IApronDeviceGroup } from "../models/apronDeviceGroup";
import { EventBus } from "../models/events";

@vts.component({ components: { } })
export default class ListDevice extends Vue {
    @vts.prop() public devices: IApronDeviceGroup[] = [];
    public deleteDevice(device: IApronDeviceGroup) {
        EventBus.$once(ACTIONS.REMOVED_DEVICE, () => {
            // console.log(`Removed device ${device.ID}`);
            device.ActionInProgress = false;
        });

        // console.log(`Removing device ${device.ID}`);
        device.ActionInProgress = true;
        EventBus.$emit(ACTIONS.REMOVE_DEVICE, device.ID);
    }

    public updateDevice(device: IApronDeviceGroup) {
        EventBus.$once(ACTIONS.UPDATED_DEVICE, () => {
            // console.log(`Updated device ${device.ID}`);
            device.ActionInProgress = false;
        });

        // console.log(`Updating device ${device.ID}`);
        device.ActionInProgress = true;
        EventBus.$emit(ACTIONS.UPDATE_DEVICE, device.ID, device.Name);
    }
}
