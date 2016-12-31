import Vue = require("vue");
import * as vts from "vue-typescript-component";
import { ACTIONS } from "../models/actions";
import { IApronDeviceGroup } from "../models/apronDeviceGroup";
import { EventBus } from "../models/events";

@vts.component({ components: { } })
export default class DeleteNode extends Vue {
    @vts.prop() public group: IApronDeviceGroup;

    public deleteNode(node: IApronDeviceGroup) {
        EventBus.$once(ACTIONS.REMOVED_DEVICE_FROM_GROUP, () => {
            // console.log(`Removed ${node.ID} from group ${this.group.ID}`);
            this.group.ActionInProgress = false;
        });

        // console.log(`Removing ${node.ID} from group ${this.group.ID}`);
        this.group.ActionInProgress = true;
        EventBus.$emit(ACTIONS.REMOVE_DEVICE_FROM_GROUP, this.group.ID, node.ID);
    }
}
