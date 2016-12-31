import Vue = require("vue");
import * as vts from "vue-typescript-component";
import { ACTIONS } from "../models/actions";
import { IApronDeviceGroup } from "../models/apronDeviceGroup";
import { EventBus } from "../models/events";

@vts.component({ components: { } })
export default class AddNode extends Vue {
    @vts.prop() public devices: IApronDeviceGroup[] = [];
    @vts.prop() public group: IApronDeviceGroup;
    public nodeId: number = 0;

    public addNode() {
        EventBus.$once(ACTIONS.ADDED_DEVICE_TO_GROUP, () => {
            // console.log(`Added ${this.nodeId} to group ${this.group.ID}`);
            this.nodeId = 0;
            this.group.ActionInProgress = false;
        });

        // console.log(`Adding ${this.nodeId} to group ${this.group.ID}`);
        this.group.ActionInProgress = true;
        EventBus.$emit(ACTIONS.ADD_DEVICE_TO_GROUP, this.group.ID, this.nodeId);
    }
}
