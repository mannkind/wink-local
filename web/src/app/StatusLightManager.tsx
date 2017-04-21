import axios from "axios";
import * as React from "react";
import { SketchPicker } from "react-color";

interface IStatusLightManagerState {
    color: IStatusLightManagerColors;
}

interface IStatusLightManagerColorsFull {
    hex: string;
    hsv: {
        h: number;
        s: number;
        v: number;
        a: number;
    };
    rgba: {
        r: number;
        g: number;
        b: number;
        a: number;
    };
    a: number;
}

type IStatusLightManagerColors = Partial<IStatusLightManagerColorsFull>;

export default class StatusLightManager extends React.Component<any, IStatusLightManagerState> {
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
        );
    }

    private handleChange = (val: IStatusLightManagerColors) => {
        if (val.hex == null) {
            return;
        }

        axios.post("/status_light/rgb/update", {
            color: val.hex,
        }).then((response) => {
           this.setState({ color: val });
        });
    }
}
