export namespace main {
	
	export class GCPProject {
	    projectId: string;
	    projectName: string;
	    isActive: boolean;
	
	    static createFrom(source: any = {}) {
	        return new GCPProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.projectId = source["projectId"];
	        this.projectName = source["projectName"];
	        this.isActive = source["isActive"];
	    }
	}
	export class PortForwardConfig {
	    localPort: number;
	    remotePort: number;
	    projectId: string;
	    zone: string;
	    instance: string;
	    status: string;
	
	    static createFrom(source: any = {}) {
	        return new PortForwardConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.localPort = source["localPort"];
	        this.remotePort = source["remotePort"];
	        this.projectId = source["projectId"];
	        this.zone = source["zone"];
	        this.instance = source["instance"];
	        this.status = source["status"];
	    }
	}

}

