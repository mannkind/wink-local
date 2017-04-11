import * as React from "react";
import Apron from "../../models/Apron";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";
import * as Actionable from "../../redux/ActionCreators";
import IAppPropsWithStore from "../../redux/State";

interface IListDeviceProps extends IAppPropsWithStore { device: IApronDeviceGroup; }
export default class ListDevice extends React.Component<IListDeviceProps, any> {
    public render() {
        const device = this.props.device;
        return (
            <div className="col col-sm-4" style={{marginBottom: "1rem"}} key={device.ID}>
                <div className="input-group input-group-sm">
                    <span className="input-group-btn">
                        <a onClick={this.removeDevice}
                            className="btn btn-sm btn-danger"
                            tabIndex={-1}>Delete</a>
                    </span>
                    <span className="input-group-addon">{device.ID}</span>
                    <input onChange={this.updateDevice}
                        className="form-control form-control-sm"
                        name="name" type="text" value={device.Name}/>
                    <span className="input-group-btn">
                        <a onClick={this.saveDevice}
                            className="btn btn-sm btn-primary"
                            tabIndex={-1}>Save</a>
                    </span>
                </div>
            </div>
        );
    }

    private removeDevice = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        Apron.removeDevice(self.props.device.ID).then(() => {
            this.props.store.dispatch(Actionable.removeDevice(self.props.device.ID));
        });
    }

    private updateDevice = (event: React.FormEvent<HTMLInputElement>) => {
        const self = this;
        const name = event.currentTarget.value;
        this.props.store.dispatch(Actionable.updateDevice(self.props.device.ID, name));
    }

    private saveDevice = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        Apron.updateDevice(self.props.device.ID, self.props.device.Name).then(() => {
            self.props.store.dispatch(Actionable.saveDevice());
        });
    }
}
