import { createStore } from "redux";
import { AppReducer } from "./Reducers";
import { IAppState } from "./State";

const AppInitialState: IAppState = {
    devices: [],
    groups: [],
};

export default createStore<IAppState>(AppReducer, AppInitialState);
