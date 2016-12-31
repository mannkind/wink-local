import Vue = require("vue");
import axios from "axios";
import {Sketch} from "vue-color";
import * as vts from "vue-typescript-component";

@vts.component({ components: { Sketch } })
export default class StatusLight extends Vue {
    public colors: IStatusLightColors = {
        a: 1,
        hex: "#000000",
    };

    public change(val: IStatusLightColors) {
        this.colors = val;
        if (this.colors.hex == null) {
            return;
        }

        axios.post("/status_light/rgb/update", {
            color: this.colors.hex,
        })
        .then((response) => {
            // console.log(response);
        })
        .catch((reason) => {
            // console.log(reason);
        });
    };
}

interface IStatusLightColors {
    hex?: string;
    hsv?: {
        h: number;
        s: number;
        v: number;
        a: number;
    };
    rgba?: {
        r: number;
        g: number;
        b: number;
        a: number;
    };
    a?: number;
}
