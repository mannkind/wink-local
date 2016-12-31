import Vue = require("vue");
import * as vts from "vue-typescript-component";

@vts.component({ components: { } })
class Events extends Vue {}

export const EventBus = new Events();
