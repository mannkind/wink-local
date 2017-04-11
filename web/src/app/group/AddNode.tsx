import * as React from "react";
import Apron from "../../models/Apron";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";
import * as Actionable from "../../redux/ActionCreators";
import IAppPropsWithStore from "../../redux/State";

interface IAddNodeProps extends IAppPropsWithStore { group: IApronDeviceGroup; devices: IApronDeviceGroup[]; }
interface IAddNodeState { nodeId: number; }
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
                <select onChange={this.handleChange} className="form-control form-control-sm">
                    <option value={this.state.nodeId}>Select a Device</option>
                    {deviceOpts}
                </select>
                <span className="input-group-btn">
                    <a tabIndex={-1} className="btn btn-sm btn-primary"
                        onClick={this.addNode}>Add</a>
                </span>
            </div>
        );
    }

    private handleChange = (event: React.FormEvent<HTMLSelectElement>) => {
        this.setState({ nodeId: +event.currentTarget.value });
    }

    private addNode = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        Apron.addDeviceToGroup(self.props.group.ID, self.state.nodeId).then(() => {
            self.props.store.dispatch(Actionable.addDeviceToGroup(self.props.group.ID, self.state.nodeId));
            self.setState({ nodeId: 0 });
        });
    }
}
