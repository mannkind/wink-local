import * as React from "react";
import Apron from "../models/Apron";
import IApronDeviceGroup from "../models/ApronDeviceGroup";
import * as Actionable from "../redux/ActionCreators";
import IAppPropsWithStore from "../redux/State";
import AddDevice from "./device/AddDevice";
import ListDevice from "./device/ListDevice";

interface IDeviceManagerProps extends IAppPropsWithStore {
    devices: IApronDeviceGroup[];
}

export default class DeviceManager extends React.Component<IDeviceManagerProps, void> {
    public render() {
        const { devices, store } = this.props;

        const existingDevices = devices.map((device, index) => {
            return (
                <ListDevice
                    key={device.ID}
                    device={device}
                    removeDevice={this.removeDevice}
                    updateDevice={this.updateDevice}
                    saveDevice={this.saveDevice} />
            );
        });

        return (
            <div className="device-container">
                <h2>Device Manager</h2>
                <div className="card">
                    <div className="card-block">
                        <div className="row">
                            <div className="col-md-4">
                                <h3>New Device</h3>
                                <AddDevice addDevice={this.addDevice} />
                            </div>
                        </div>
                        <hr />
                        <div className="row">
                            <div className="col-md-12">
                                <h3>Existing Devices</h3>
                                <div className="row">
                                    {existingDevices}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    private addDevice = (radio: string) => {
        const self = this;
        const timeout = 60000;
        Apron.addDevice(radio).then(() => {
            setTimeout(() => {
                Apron.listDevices().then((response) => {
                    self.props.store.dispatch(Actionable.addDevice(response.data));
                });
            }, timeout);
        });
    }

    private removeDevice = (deviceId: number) => {
        const self = this;
        Apron.removeDevice(deviceId).then(() => {
            this.props.store.dispatch(Actionable.removeDevice(deviceId));
        });
    }

    private updateDevice = (deviceId: number, name: string) => {
        this.props.store.dispatch(Actionable.updateDevice(deviceId, name));
    }

    private saveDevice = (deviceId: number, name: string) => {
        const self = this;
        Apron.updateDevice(deviceId, name).then(() => {
            self.props.store.dispatch(Actionable.saveDevice());
        });
    }
}
