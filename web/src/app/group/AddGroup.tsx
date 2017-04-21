import * as React from "react";
import IApronDeviceGroup from "../../models/ApronDeviceGroup";

interface IAddGroupProps {
    addGroup(name: string);
}

interface IAddGroupStateFull {
    name: string;
}

type IAddGroupState = Partial<IAddGroupStateFull>;

export default class AddGroup extends React.Component<IAddGroupProps, IAddGroupState> {
    constructor(props: IAddGroupProps, context) {
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
                    onChange={this.onChange} value={this.state.name} />
                <span className="input-group-btn">
                    <a className="btn btn-sm btn-primary"
                        onClick={this.onClick}
                        tabIndex={-1} >Add</a>
                </span>
            </div>
        );
    }

    private onChange = (event: React.FormEvent<HTMLInputElement>) => {
        this.setState({ name: event.currentTarget.value });
    }

    private onClick = (event: React.FormEvent<HTMLAnchorElement>) => {
        event.preventDefault();

        this.props.addGroup(this.state.name);
        this.setState({ name: "" });
    }
}
