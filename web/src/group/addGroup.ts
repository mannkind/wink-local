import Vue = require("vue");
import * as vts from "vue-typescript-component";
import { ACTIONS } from "../models/actions";
import { EventBus } from "../models/events";

@vts.component({ components: { } })
export default class AddGroup extends Vue {
    public name: string = "";
    public actionInProgress: boolean = false;

    public addGroup() {
        EventBus.$once(ACTIONS.ADDED_GROUP, () => {
            // console.log(`Group ${this.name} added`);
            this.actionInProgress = false;
            this.name = "";
        });

        // console.log(`Adding group ${this.name}`);
        this.actionInProgress = true;
        EventBus.$emit(ACTIONS.ADD_GROUP, this.name);
    }
}
