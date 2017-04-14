import * as React from "react";
import Apron from "../../models/Apron";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";
import * as Actionable from "../../redux/ActionCreators";
import IAppPropsWithStore from "../../redux/State";
import AddNode from "./AddNode";
import DeleteNode from "./DeleteNode";

interface IListGroupProps extends IAppPropsWithStore { devices: IApronDeviceGroup[]; group: IApronDeviceGroup; }
export default class ListGroup extends React.Component<IListGroupProps, void> {
    public render() {
        const { devices, group, store } = this.props;
        const groupNodes = group.Nodes.map((node, index1) => {
            return (<DeleteNode store={store} group={group} node={node} />);
        });
        return (
            <div>
                <div className="row" style={{marginBottom: "1rem"}}>
                    <div className="col-sm-4">
                        <div className="input-group input-group-sm">
                            <span className="input-group-btn">
                                <a onClick={this.removeGroup}
                                    tabIndex={-1} className="btn btn-sm btn-danger">Delete</a>
                            </span>
                            <span className="input-group-addon">{group.ID}</span>
                            <input value={group.Name}
                                className="form-control form-control-sm"
                                readOnly={true} type="text" />
                            <span className="input-group-addon">
                                {group.Nodes.length}</span>
                        </div>
                    </div>
                    <div className="col-sm-8">
                        <AddNode store={store} group={group} devices={devices} />
                    </div>
                </div>
                <div className="row">
                    <div className="col-sm-2">
                        <strong>Assoc. Devices</strong>
                    </div>
                    <div className="col-sm-10">
                        <div className="row">
                            {groupNodes}
                        </div>
                    </div>
                </div>
            </div>
        );
    }

    private removeGroup = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        Apron.removeGroup(self.props.group.ID).then(() => {
            self.props.store.dispatch(Actionable.removeGroup(self.props.group.ID));
        });
    }
}
