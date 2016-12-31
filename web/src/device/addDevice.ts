import Vue = require("vue");
import * as vts from "vue-typescript-component";
import { ACTIONS } from "../models/actions";
import { EventBus } from "../models/events";

@vts.component({ components: { } })
export default class AddDevice extends Vue {
    public availableRadios: string[] = [
        "Zigbee",
        "ZWave",
        "Lutron",
    ];

    public radio: string = "";
    public progress: number = 0;
    private progressTimeout: number = 1000;

    public addDevice() {
        if (this.radio === "") {
            return;
        }

        // Initiate pairing
        // console.log(`Pairing initiated for ${this.radio}`);
        EventBus.$emit(ACTIONS.ADD_DEVICE, this.radio);

        // Update the progress bar
        const interval = setInterval(() => {
            // console.log(`Progress updated for ${this.radio}`);
            this.progress++;
        }, this.progressTimeout);

        // When the device is added, stop the progress bar and reset the radio selection
        EventBus.$once(ACTIONS.ADDED_DEVICE, () => {
            // console.log(`Pairing stopped for ${this.radio}`);
            clearInterval(interval);

            this.radio = "";
            this.progress = 0;
        });
    };
}
