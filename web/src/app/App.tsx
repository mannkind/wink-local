import axios from "axios";
import * as React from "react";
import { Unsubscribe } from "redux";
import * as Actionable from "../redux/ActionCreators";
import IAppPropsWithStore from "../redux/State";
import AddDevice from "./device/AddDevice";
import ListDevice from "./device/ListDevice";
import AddGroup from "./group/AddGroup";
import ListGroup from "./group/ListGroup";
import StatusLight from "./StatusLight";

export default class App extends React.Component<IAppPropsWithStore, void> {
    private unsubscribe: Unsubscribe;

    public componentDidMount() {
        const self = this;
        self.unsubscribe = this.props.store.subscribe(() => {

            self.forceUpdate();
        });

        this.updateDeviceList();
        this.updateGroupList();
    }

    public componentWillUnmount() {
        this.unsubscribe();
    }

    public render() {
        const title = "Wink-Local UI";
        const store = this.props.store;
        const state = store.getState();
        const devices = state.devices;
        const groups = state.groups;

        const existingDevices = devices.map((device, index) => {
            return (<ListDevice store={store} device={device} key={device.ID} />);
        });

        const existingGroups = groups.map((group, index) => {
            return (<ListGroup store={store} group={group} devices={devices} key={group.ID} />);
        });

        return (
            <div className="container">
                <h1>{title}</h1>
                <div className="row">
                    <div className="col-md-12">
                        <h2>Device Manager</h2>
                        <div className="card">
                            <div className="card-block">
                                <div className="row">
                                    <div className="col-md-4">
                                        <h3>New Device</h3>
                                        <AddDevice store={store} />
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
                        <h2>Group Manager</h2>
                        <div className="card">
                            <div className="card-block">
                                <div className="row">
                                    <div className="col-md-4">
                                        <h3>New Group</h3>
                                        <AddGroup store={store} />
                                    </div>
                                </div>
                                <hr />
                                <h3>Existing Groups</h3>
                                {existingGroups}
                            </div>
                        </div>
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-12">
                        <StatusLight />
                    </div>
                </div>
            </div>
        );
    }

    private updateDeviceList = (): Promise<void> => {
        const self = this;
        return axios
            .get("/device/list")
            .then((response) => {
                self.props.store.dispatch(Actionable.initDevices(response.data));
            });
    }

    private updateGroupList = (): Promise<void> => {
        const self = this;
        return axios
            .get("/group/list")
             .then((response) => {
                self.props.store.dispatch(Actionable.initGroups(response.data));
            });
    }
}
