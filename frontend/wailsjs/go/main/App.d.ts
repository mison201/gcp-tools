// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function AuthenticateGCloud():Promise<void>;

export function CheckGCloudAuth():Promise<boolean>;

export function GetActiveProject():Promise<string>;

export function GetComputeInstances(arg1:string,arg2:string):Promise<Array<string>>;

export function GetGCPProjects():Promise<Array<main.GCPProject>>;

export function GetPortForwardStatus(arg1:number):Promise<string>;

export function GetZones(arg1:string):Promise<Array<string>>;

export function ListPortForwards():Promise<Array<main.PortForwardConfig>>;

export function SetActiveProject(arg1:string):Promise<void>;

export function StartPortForward(arg1:main.PortForwardConfig):Promise<void>;

export function StopPortForward(arg1:number):Promise<void>;

export function TestConnection(arg1:number):Promise<boolean>;
