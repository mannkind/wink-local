import * as React from "react";
import Apron from "../../models/Apron";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";
import * as Actionable from "../../redux/ActionCreators";
import IAppPropsWithStore from "../../redux/State";

interface IDeleteNodeProps extends IAppPropsWithStore { group: IApronDeviceGroup; node: IApronDeviceGroup; }
export default class DeleteNode extends React.Component<IDeleteNodeProps, any> {
    public render() {
        const group = this.props.group;
        const node = this.props.node;
        return (
            <div className="col col-md-4" style={{marginBottom: "1rem"}} key={node.ID}>
                <div className="input-group input-group-sm">
                    <span className="input-group-btn">
                        <a tabIndex={-1} className="btn btn-sm btn-danger"
                            onClick={this.deleteNode}>Delete</a>
                    </span>
                    <input value={node.Name}
                        className="form-control form-control-sm"
                        readOnly={true} type="text" />
                </div>
            </div>
        );
    }

    private deleteNode = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        Apron.removeDeviceFromGroup(self.props.group.ID, self.props.node.ID).then((response) => {
            self.props.store.dispatch(Actionable.removeDeviceFromGroup(this.props.group.ID, self.props.node.ID));
        });
    }
}
