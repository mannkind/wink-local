import axios from "axios";
import * as React from "react";
import { SketchPicker } from "react-color";

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

interface IStatusLightState { color: IStatusLightColors; }
export default class StatusLight extends React.Component<any, IStatusLightState> {
    constructor(props, context) {
        super(props, context);

        this.state = {
            color: {
                a: 1,
                hex: "#000000",
            },
        };
    }

    public render() {
        return (
            <div>
                <div className="row">
                    <div className="col-md-12">
                        <h2>Status Light Manager</h2>
                        <div className="card">
                            <div className="card-block">
                                <SketchPicker color={this.state.color}
                                    onChangeComplete={this.handleChange} />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    private handleChange = (val: IStatusLightColors) => {
        if (val.hex == null) {
            return;
        }

        axios.post("/status_light/rgb/update", {
            color: val.hex,
        })
        .then((response) => {
           this.setState({ color: val });
        });
    }
}
