import axios from "axios";
import { createStore } from "redux";
import { IAppActions, IAppActionTypes } from "./Actions";
import { AppReducer } from "./Reducers";
import { IAppState } from "./State";

const AppInitialState: IAppState = {
    devices: [],
    groups: [],
};

export default createStore<IAppState>(AppReducer, AppInitialState);
