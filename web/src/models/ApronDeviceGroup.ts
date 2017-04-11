interface IApronDeviceGroup {
    ID: number;
    Name: string;
    Nodes?: IApronDeviceGroup[];
}

export default IApronDeviceGroup;
