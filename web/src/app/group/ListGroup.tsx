import * as React from "react";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";
import AddNode from "./AddNode";
import DeleteNode from "./DeleteNode";

interface IListGroupProps  {
    devices: IApronDeviceGroup[];
    group: IApronDeviceGroup;
    deleteGroup(groupId: number);
    addNode(groupId: number, nodeId: number);
    deleteNode(groupId: number, nodeId: number);
}

export default class ListGroup extends React.Component<IListGroupProps, void> {
    public render() {
        const { devices, group } = this.props;
        const groupNodes = group.Nodes.map((node, index1) => {
            return (<DeleteNode group={group} node={node} deleteNode={this.props.deleteNode}/>);
        });
        return (
            <div>
                <div className="row" style={{marginBottom: "1rem"}}>
                    <div className="col-sm-4">
                        <div className="input-group input-group-sm">
                            <span className="input-group-btn">
                                <a onClick={this.onClick}
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
                        <AddNode group={group} devices={devices} addNode={this.props.addNode} />
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

    private onClick = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        this.props.deleteGroup(this.props.group.ID);
    }
}
