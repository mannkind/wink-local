import * as React from "react";

interface IAddDeviceProps {
    addDevice(radio: string);
}

interface IAddDeviceStateFull {
    radio: string;
}

type IAddDeviceState = Partial<IAddDeviceStateFull>;

export default class AddDevice extends React.Component<IAddDeviceProps, IAddDeviceState> {
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
                    <select onChange={this.onChange}
                        className="form-control form-control-sm">
                        <option value={this.state.radio}>Select a Radio</option>
                        {radioOpts}
                    </select>
                    <span className="input-group-btn">
                        <a onClick={this.onClick}
                            className="btn btn-sm btn-primary"
                            tabIndex={-1}>Add</a>
                    </span>
                </div>
            </div>
        );
    }

    private onChange = (event: React.FormEvent<HTMLSelectElement>) => {
        this.setState({ radio: event.currentTarget.value});
    }

    private onClick = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        this.props.addDevice(this.state.radio);
        this.setState({ radio: "" });
    }
}
