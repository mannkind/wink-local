import axios, { Promise } from "axios";

class ApronService {
    public addDevice(radio: string): Promise<void> {
        return axios.post("/device/add", {
            Radio: radio.toLowerCase(),
        })
        .catch((reason) => {
            // console.log(`Failed to addDevice because ${reason}`);
        });
    };

    public removeDevice(deviceId: number): Promise<void> {
        return axios.post("/device/delete", {
            ID: deviceId.toString(),
        })
        .catch((reason) => {
            // console.log(`Failed to deleteDevice because ${reason}`);
        });
    }

    public updateDevice(deviceId: number, name: string): Promise<void> {
        return axios.post(`/device/${deviceId}/update`, {
            Name: name,
        })
        .catch((reason) => {
            // console.log(`Failed to updateDevice because ${reason}`);
        });
    }

    public addGroup(name: string): Promise<void> {
        return axios.post("/group/add", {
            Name: name,
        })
        .catch((reason) => {
            // console.log(`Failed to addGroup because ${reason}`);
        });
    }

    public removeGroup(groupId: number): Promise<void> {
        return axios.post("/group/delete", {
            ID: groupId.toString(),
        })
        .catch((reason) => {
            // console.log(`Failed to deleteGroup because ${reason}`);
        });
    }

    public addDeviceToGroup(groupId: number, nodeId: number): Promise<void> {
        return axios.post(`/group/${groupId}/add`, {
            DeviceID: nodeId.toString(),
        })
        .catch((reason) => {
            // console.log(`Failed to addNode because ${reason}`); 
        });
    }

    public removeDeviceFromGroup(groupId: number, deviceId: number): Promise<void> {
        return axios.post(`/group/${groupId}/delete`, {
            DeviceID: deviceId.toString(),
        })
        .catch((reason) => {
            // console.log(`Failed to deleteNode because ${reason}`); 
        });
    }
}

export const Apron = new ApronService();
