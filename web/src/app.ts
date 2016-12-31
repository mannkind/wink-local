import Vue = require("vue");
import * as vts from "vue-typescript-component";
import * as DeviceGroup from "./deviceGroup.vue";
import * as StatusLight from "./statusLight.vue";

@vts.component({ components: { DeviceGroup, StatusLight } })
export default class App extends Vue {
    public title: string = "Wink-Local UI";
}
