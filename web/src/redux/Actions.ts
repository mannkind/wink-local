import IApronDeviceGroup from "../models/ApronDeviceGroup";

export const enum IAppActionTypes {
    INIT_DEVICES,
    INIT_GROUPS,
    ADD_DEVICE,
    ADD_DEVICE_TO_GROUP,
    UPDATE_DEVICE,
    REMOVE_DEVICE,
    SAVE_DEVICE,
    ADD_GROUP,
    REMOVE_GROUP,
    REMOVE_DEVICE_FROM_GROUP,
}

export type IAppActions = {
    type: IAppActionTypes.INIT_DEVICES;
    payload: IApronDeviceGroup[];
} | {
    type: IAppActionTypes.INIT_GROUPS;
    payload: IApronDeviceGroup[];
} | {
    type: IAppActionTypes.ADD_DEVICE;
    payload: IApronDeviceGroup[];
}  | {
    type: IAppActionTypes.ADD_DEVICE_TO_GROUP;
    payload: { groupId: number, deviceId: number };
}  | {
    type: IAppActionTypes.REMOVE_DEVICE_FROM_GROUP;
    payload: { groupId: number, deviceId: number };
} | {
    type: IAppActionTypes.UPDATE_DEVICE;
    payload: { deviceId: number, name: string };
} | {
    type: IAppActionTypes.REMOVE_DEVICE;
    payload: number;
} | {
    type: IAppActionTypes.SAVE_DEVICE;
    payload: null;
} | {
    type: IAppActionTypes.ADD_GROUP;
    payload: IApronDeviceGroup[];
} | {
    type: IAppActionTypes.REMOVE_GROUP;
    payload: number;
};
