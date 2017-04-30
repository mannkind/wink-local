import * as React from "react";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";

interface IDeleteNodeProps {
    group: IApronDeviceGroup;
    node: IApronDeviceGroup;
    deleteNode(groupId: number, nodeId: number);
}

export default class DeleteNode extends React.Component<IDeleteNodeProps, any> {
    public render() {
        const { node } = this.props;
        return (
            <div className="col col-md-4" style={{marginBottom: "1rem"}} key={node.ID}>
                <div className="input-group input-group-sm">
                    <span className="input-group-btn">
                        <a tabIndex={-1} className="btn btn-sm btn-danger"
                            onClick={this.onClick}>Delete</a>
                    </span>
                    <input value={node.Name}
                        className="form-control form-control-sm"
                        readOnly={true} type="text" />
                </div>
            </div>
        );
    }

    private onClick = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();
        this.props.deleteNode(this.props.group.ID, this.props.node.ID);
    }
}
