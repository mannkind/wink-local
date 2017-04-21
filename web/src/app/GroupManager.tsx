import * as React from "react";
import Apron from "../models/Apron";
import IApronDeviceGroup from "../models/ApronDeviceGroup";
import * as Actionable from "../redux/ActionCreators";
import IAppPropsWithStore from "../redux/State";
import AddGroup from "./group/AddGroup";
import ListGroup from "./group/ListGroup";

interface IGroupManagerProps extends IAppPropsWithStore {
    devices: IApronDeviceGroup[];
    groups: IApronDeviceGroup[];
}

export default class GroupManager extends React.Component<IGroupManagerProps, void> {

    public render() {
        const { devices, groups, store } = this.props;

        const existingGroups = groups.map((group, index) => {
            return (
                <ListGroup
                    key={group.ID}
                    devices={devices}
                    group={group}
                    addNode={this.addNode}
                    deleteNode={this.deleteNode}
                    deleteGroup={this.deleteGroup} />);
        });

        return (
            <div className="group_container">
                <h2>Group Manager</h2>
                <div className="card">
                    <div className="card-block">
                        <div className="row">
                            <div className="col-md-4">
                                <h3>New Group</h3>
                                <AddGroup addGroup={this.addGroup} />
                            </div>
                        </div>
                        <hr />
                        <h3>Existing Groups</h3>
                        {existingGroups}
                    </div>
                </div>
            </div>
        );
    }

    private addGroup = (name: string) => {
        const self = this;
        Apron.addGroup(name).then(() => {
            Apron.listGroups().then((response) => {
                self.props.store.dispatch(Actionable.addGroup(response.data));
            });
        });
    }

    private addNode = (groupId: number, nodeId: number) => {
        const self = this;
        Apron.addDeviceToGroup(groupId, nodeId).then(() => {
            self.props.store.dispatch(Actionable.addDeviceToGroup(groupId, nodeId));
        });
    }

    private deleteGroup = (groupId: number) => {
        const self = this;
        Apron.removeGroup(groupId).then(() => {
            self.props.store.dispatch(Actionable.removeGroup(groupId));
        });
    }

    private deleteNode = (groupId: number, nodeId: number) => {
        const self = this;
        Apron.removeDeviceFromGroup(groupId, nodeId).then((response) => {
            self.props.store.dispatch(Actionable.removeDeviceFromGroup(groupId, nodeId));
        });
    }
}
