import * as React from "react";
import { Unsubscribe } from "redux";
import Apron from "../models/Apron";
import * as Actionable from "../redux/ActionCreators";
import IAppPropsWithStore from "../redux/State";
import DeviceManager from "./DeviceManager";
import GroupManager from "./GroupManager";
import StatusLightManager from "./StatusLightManager";

export default class App extends React.Component<IAppPropsWithStore, void> {
    private unsubscribe: Unsubscribe;

    public componentDidMount() {
        const self = this;
        self.unsubscribe = this.props.store.subscribe(() => {

            self.forceUpdate();
        });

        Apron.listDevices().then((response) => {
            self.props.store.dispatch(Actionable.initDevices(response.data));
        });

        Apron.listGroups().then((response) => {
            self.props.store.dispatch(Actionable.initGroups(response.data));
        });
    }

    public componentWillUnmount() {
        this.unsubscribe();
    }

    public render() {
        const title = "Wink-Local UI";
        const { store } = this.props;
        const { devices, groups } = store.getState();

        return (
            <div className="container">
                <h1>{title}</h1>
                <div className="row">
                    <div className="col-md-12">
                        <DeviceManager store={store} devices={devices} />
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-12">
                        <GroupManager store={store} devices={devices} groups={groups} />
                    </div>
                </div>
                <div className="row">
                    <div className="col-md-12">
                        <StatusLightManager />
                    </div>
                </div>
            </div>
        );
    }
}
