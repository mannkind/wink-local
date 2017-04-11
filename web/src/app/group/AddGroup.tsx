import axios from "axios";
import * as React from "react";
import Apron from "../../models/Apron";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";
import * as Actionable from "../../redux/ActionCreators";
import IAppPropsWithStore from "../../redux/State";

interface IAddGroupStateFull { name: string; }
type IAddGroupState = Partial<IAddGroupStateFull>;
export default class AddGroup extends React.Component<IAppPropsWithStore, IAddGroupState> {
    constructor(props: IAppPropsWithStore, context) {
        super(props, context);

        this.state = {
            name: "",
        };
    }

    public render() {
        return (
            <div className="input-group input-group-sm">
                <input className="form-control form-control-sm" type="text"
                    placeholder="Enter a group name"
                    onChange={this.handleChange} value={this.state.name} />
                <span className="input-group-btn">
                    <a className="btn btn-sm btn-primary"
                        onClick={this.addGroup}
                        tabIndex={-1} >Add</a>
                </span>
            </div>
        );
    }

    private handleChange = (event: React.FormEvent<HTMLInputElement>) => {
        this.setState({ name: event.currentTarget.value });
    }

    private addGroup = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        const self = this;
        Apron.addGroup(self.state.name).then(() => {
            axios
                .get("/group/list")
                .then((response) => {
                    self.props.store.dispatch(Actionable.addGroup(response.data));
                    this.setState({ name: "" });
                });
        });
    }
}
