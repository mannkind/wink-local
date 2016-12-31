export interface IApronDeviceGroup {
    ActionInProgress: boolean;
    ID: number;
    Name: string;
    Nodes?: IApronDeviceGroup[];
}
