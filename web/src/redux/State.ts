import { Store } from "redux";
import IApronDeviceGroup from "../models/ApronDeviceGroup";

interface IAppStateFull {
    devices: IApronDeviceGroup[];
    groups: IApronDeviceGroup[];
}

export type IAppState = Partial<IAppStateFull>;

interface IAppPropsWithStore {
    store: Store<IAppState>;
}

export default IAppPropsWithStore;
