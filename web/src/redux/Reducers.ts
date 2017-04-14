import IApronDeviceGroup from "../models/ApronDeviceGroup";
import { IAppActions, IAppActionTypes } from "./Actions";
import { IAppState } from "./State";

/* tslint:disable:no-namespace */
declare global {
    /* tslint:disable:interface-name */
    interface Array<T> {
        find(predicate: (value: T, index: number, obj: T[]) => boolean, thisArg?: T): T;
        findIndex(predicate: (value: T, index: number, obj: T[]) => boolean, thisArg?: T): number;
    }
}

export const AppReducer = (state: IAppState, action: IAppActions): IAppState => {
        switch (action.type) {
            case IAppActionTypes.INIT_DEVICES:
                return { ...state, ...{ devices: action.payload }};

            case IAppActionTypes.INIT_GROUPS:
                return { ...state, ...{ groups: action.payload }};

            case IAppActionTypes.ADD_DEVICE:
                return { ...state, ...{ devices: action.payload }};

            case IAppActionTypes.REMOVE_DEVICE:
                return (() => {
                    const devices = state.devices.filter((device) => device.ID !== action.payload);
                    const { groups } = state;

                    groups.forEach((x) => {
                        if (x.Nodes == null) {
                            return;
                        }

                        const index = x.Nodes.findIndex((y) => y.ID === action.payload);
                        if (index === -1) {
                            return;
                        }

                        x.Nodes.splice(index, 1);
                    });

                    return { ...state, ...{ devices, groups }};
                })();

            case IAppActionTypes.UPDATE_DEVICE:
                return (() => {
                    const { devices, groups } = state;
                    const device = devices.find((x) => x.ID === action.payload.deviceId);
                    device.Name = action.payload.name;

                    groups.forEach((x) => {
                        if (x.Nodes == null) {
                            return;
                        }

                        const groupDevice = x.Nodes.find((x1) => x1.ID === device.ID);
                        if (groupDevice == null) {
                            return;
                        }

                        groupDevice.Name = action.payload.name;
                    });

                    return {...state, ...{ devices, groups }};
                })();

            case IAppActionTypes.SAVE_DEVICE:
                return state;

            case IAppActionTypes.ADD_DEVICE_TO_GROUP:
                return (() => {
                    const { devices, groups } = state;
                    const device = devices.find((x) => x.ID === action.payload.deviceId);
                    const group = groups.find((x) => x.ID === action.payload.groupId);

                    if (group == null || group.Nodes == null || device == null) {
                        return state;
                    }

                    group.Nodes.push(device);

                    return {...state, ...{ devices, groups }};
                })();

            case IAppActionTypes.REMOVE_DEVICE_FROM_GROUP:
                return (() => {
                    const { devices, groups } = state;
                    const device = devices.find((x) => x.ID === action.payload.deviceId);
                    const group = groups.find((x) => x.ID === action.payload.groupId);

                    if (group == null || group.Nodes == null) {
                        return state;
                    }

                    group.Nodes = group.Nodes.filter((node) => node.ID !== action.payload.deviceId);

                    return {...state, ...{ devices, groups }};
                })();

            case IAppActionTypes.ADD_GROUP:
                return { ...state, ...{ groups: action.payload }};

            case IAppActionTypes.REMOVE_GROUP:
                return (() => {
                    const groups = state.groups.filter((group) => group.ID !== action.payload);

                    return {...state, ...{ groups }};
                })();

            default:
                return state;
        }
    };
