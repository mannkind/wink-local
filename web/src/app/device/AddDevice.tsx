import axios from "axios";
import * as React from "react";
import Apron from "../../models/Apron";
import * as Actionable from "../../redux/ActionCreators";
import IAppPropsWithStore from "../../redux/State";

interface IAddDeviceStateFull { radio: string; }
type IAddDeviceState = Partial<IAddDeviceStateFull>;
export default class AddDevice extends React.Component<IAppPropsWithStore, IAddDeviceState> {
    public availableRadios: string[] = [
        "Zigbee",
        "ZWave",
        "Lutron",
    ];

    constructor(props, context) {
        super(props, context);

        this.state = {
            radio: "",
        };
    }

    public render() {
        const radioOpts = this.availableRadios.map((radio, index) => {
            return (<option value={radio} key={radio}>{radio}</option>);
        });

        return (
            <div>
                <div className="input-group input-group-sm">
                    <select onChange={this.handleChange}
                        className="form-control form-control-sm">
                        <option value={this.state.radio}>Select a Radio</option>
                        {radioOpts}
                    </select>
                    <span className="input-group-btn">
                        <a onClick={this.addDevice}
                            className="btn btn-sm btn-primary"
                            tabIndex={-1}>Add</a>
                    </span>
                </div>
            </div>
        );
    }

    private handleChange = (event: React.FormEvent<HTMLSelectElement>) => {
        this.setState({ radio: event.currentTarget.value});
    }

    private addDevice = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        const timeout = 60000;
        Apron.addDevice(this.state.radio).then(() => {
            setTimeout(() => {
                axios
                    .get("/device/list")
                    .then((response) => {
                        self.props.store.dispatch(Actionable.addDevice(response.data));
                        self.setState({ radio: "" });
                    });
            }, timeout);
        });
    }
}
