import IApronDeviceGroup from "../models/ApronDeviceGroup";
import { IAppActions, IAppActionTypes } from "./Actions";

export const initDevices = (devices: IApronDeviceGroup[]): IAppActions => {
    return {
        payload: devices,
        type: IAppActionTypes.INIT_DEVICES,
    };
};

export const initGroups = (groups: IApronDeviceGroup[]): IAppActions => {
    return {
        payload: groups,
        type: IAppActionTypes.INIT_GROUPS,
    };
};

export const addDevice = (devices: IApronDeviceGroup[]): IAppActions => {
    return {
        payload: devices,
        type: IAppActionTypes.ADD_DEVICE,
    };
};

export const addDeviceToGroup = ( groupId: number, deviceId: number ): IAppActions => {
    return {
        payload: { groupId, deviceId },
        type: IAppActionTypes.ADD_DEVICE_TO_GROUP,
    };
};

export const removeDeviceFromGroup = ( groupId: number, deviceId: number ): IAppActions => {
    return {
        payload: { groupId, deviceId },
        type: IAppActionTypes.REMOVE_DEVICE_FROM_GROUP,
    };
};

export const updateDevice = (deviceId: number, name: string): IAppActions => {
    return {
        payload: { deviceId, name },
        type: IAppActionTypes.UPDATE_DEVICE,
    };
};

export const removeDevice = (deviceId: number): IAppActions => {
    return {
        payload: deviceId,
        type: IAppActionTypes.REMOVE_DEVICE,
    };
};

export const saveDevice = (): IAppActions => {
    return {
        payload: null,
        type: IAppActionTypes.SAVE_DEVICE,
    };
};

export const addGroup = (groups: IApronDeviceGroup[]): IAppActions => {
    return {
        payload: groups,
        type: IAppActionTypes.ADD_GROUP,
    };
};

export const removeGroup = (groupId: number): IAppActions => {
    return {
        payload: groupId,
        type: IAppActionTypes.REMOVE_GROUP,
    };
};
