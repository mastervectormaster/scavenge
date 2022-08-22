/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "mastervectormaster.scavenge.scavenge";

export interface MsgSubmitScavenge {
  creator: string;
  solutionHash: string;
  description: string;
  reward: string;
}

export interface MsgSubmitScavengeResponse {}

const baseMsgSubmitScavenge: object = {
  creator: "",
  solutionHash: "",
  description: "",
  reward: "",
};

export const MsgSubmitScavenge = {
  encode(message: MsgSubmitScavenge, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.solutionHash !== "") {
      writer.uint32(18).string(message.solutionHash);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.reward !== "") {
      writer.uint32(34).string(message.reward);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSubmitScavenge {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSubmitScavenge } as MsgSubmitScavenge;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.solutionHash = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.reward = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitScavenge {
    const message = { ...baseMsgSubmitScavenge } as MsgSubmitScavenge;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.solutionHash !== undefined && object.solutionHash !== null) {
      message.solutionHash = String(object.solutionHash);
    } else {
      message.solutionHash = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = String(object.reward);
    } else {
      message.reward = "";
    }
    return message;
  },

  toJSON(message: MsgSubmitScavenge): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.solutionHash !== undefined &&
      (obj.solutionHash = message.solutionHash);
    message.description !== undefined &&
      (obj.description = message.description);
    message.reward !== undefined && (obj.reward = message.reward);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSubmitScavenge>): MsgSubmitScavenge {
    const message = { ...baseMsgSubmitScavenge } as MsgSubmitScavenge;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.solutionHash !== undefined && object.solutionHash !== null) {
      message.solutionHash = object.solutionHash;
    } else {
      message.solutionHash = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.reward !== undefined && object.reward !== null) {
      message.reward = object.reward;
    } else {
      message.reward = "";
    }
    return message;
  },
};

const baseMsgSubmitScavengeResponse: object = {};

export const MsgSubmitScavengeResponse = {
  encode(
    _: MsgSubmitScavengeResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSubmitScavengeResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSubmitScavengeResponse,
    } as MsgSubmitScavengeResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitScavengeResponse {
    const message = {
      ...baseMsgSubmitScavengeResponse,
    } as MsgSubmitScavengeResponse;
    return message;
  },

  toJSON(_: MsgSubmitScavengeResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSubmitScavengeResponse>
  ): MsgSubmitScavengeResponse {
    const message = {
      ...baseMsgSubmitScavengeResponse,
    } as MsgSubmitScavengeResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitScavenge(
    request: MsgSubmitScavenge
  ): Promise<MsgSubmitScavengeResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SubmitScavenge(
    request: MsgSubmitScavenge
  ): Promise<MsgSubmitScavengeResponse> {
    const data = MsgSubmitScavenge.encode(request).finish();
    const promise = this.rpc.request(
      "mastervectormaster.scavenge.scavenge.Msg",
      "SubmitScavenge",
      data
    );
    return promise.then((data) =>
      MsgSubmitScavengeResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
