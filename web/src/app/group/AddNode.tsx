import * as React from "react";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";

interface IAddNodeProps {
    devices: IApronDeviceGroup[];
    group: IApronDeviceGroup;
    addNode(groupId: number, nodeId: number);
}

interface IAddNodeState {
    nodeId: number;
}

export default class AddNode extends React.Component<IAddNodeProps, IAddNodeState> {
    constructor(props: IAddNodeProps, context) {
        super(props, context);

        this.state = {
            nodeId: 0,
        };
    }

    public render() {
        const deviceOpts = this.props.devices.map((device, index) => {
            return (<option value={device.ID}>{device.Name}</option>);
        });

        return (
            <div className="input-group input-group-sm">
                <select onChange={this.onChange} className="form-control form-control-sm">
                    <option value={this.state.nodeId}>Select a Device</option>
                    {deviceOpts}
                </select>
                <span className="input-group-btn">
                    <a tabIndex={-1} className="btn btn-sm btn-primary"
                        onClick={this.onClick}>Add</a>
                </span>
            </div>
        );
    }

    private onChange = (event: React.FormEvent<HTMLSelectElement>) => {
        this.setState({ nodeId: +event.currentTarget.value });
    }

    private onClick = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        this.props.addNode(this.props.group.ID, this.state.nodeId);
        this.setState({ nodeId: 0 });
    }
}
