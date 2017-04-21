import * as React from "react";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";

interface IListDeviceProps {
    device: IApronDeviceGroup;
    removeDevice(deviceId: number);
    saveDevice(deviceId: number, name: string);
    updateDevice(deviceId: number, name: string);
}

export default class ListDevice extends React.Component<IListDeviceProps, any> {
    public render() {
        const { device } = this.props;
        return (
            <div className="col col-sm-4" style={{marginBottom: "1rem"}} key={device.ID}>
                <div className="input-group input-group-sm">
                    <span className="input-group-btn">
                        <a onClick={this.onClickDelete}
                            className="btn btn-sm btn-danger"
                            tabIndex={-1}>Delete</a>
                    </span>
                    <span className="input-group-addon">{device.ID}</span>
                    <input onChange={this.onChange}
                        className="form-control form-control-sm"
                        name="name" type="text" value={device.Name}/>
                    <span className="input-group-btn">
                        <a onClick={this.onClickSave}
                            className="btn btn-sm btn-primary"
                            tabIndex={-1}>Save</a>
                    </span>
                </div>
            </div>
        );
    }

    private onClickDelete = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        this.props.removeDevice(this.props.device.ID);
    }

    private onChange = (event: React.FormEvent<HTMLInputElement>) => {
        const name = event.currentTarget.value;
        this.props.updateDevice(this.props.device.ID, name);
    }

    private onClickSave = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        this.props.saveDevice(this.props.device.ID, this.props.device.Name);
    }
}
