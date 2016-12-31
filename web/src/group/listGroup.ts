import Vue = require("vue");
import * as vts from "vue-typescript-component";
import { ACTIONS } from "../models/actions";
import { IApronDeviceGroup } from "../models/apronDeviceGroup";
import { EventBus } from "../models/events";
import * as AddNode from "./addNode.vue";
import * as DeleteNode from "./deleteNode.vue";

@vts.component({ components: { AddNode, DeleteNode } })
export default class ListGroup extends Vue {
    @vts.prop() public devices: IApronDeviceGroup[] = [];
    @vts.prop() public groups: IApronDeviceGroup[] = [];

    public deleteGroup(group: IApronDeviceGroup) {
        EventBus.$once(ACTIONS.REMOVED_GROUP, () => {
            // console.log(`Removed group ${group.Name}`);
            group.ActionInProgress = false;
        });

        // console.log(`Removing group ${group.Name}`);
        group.ActionInProgress = true;
        EventBus.$emit(ACTIONS.REMOVE_GROUP, group.ID);
    }
}
