export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Date: any;
  Upload: any;
  WhtID: any;
};


export type Mutation = {
  __typename?: 'Mutation';
  noop?: Maybe<NoopPayload>;
  /** 「今日こと」を登録 */
  createWht?: Maybe<MutationResponse>;
};


export type MutationNoopArgs = {
  input?: Maybe<NoopInput>;
};


export type MutationCreateWhtArgs = {
  wht: WhtInput;
};

export type MutationResponse = {
  __typename?: 'MutationResponse';
  id?: Maybe<Scalars['ID']>;
};

export type Node = {
  id: Scalars['ID'];
};

export type NoopInput = {
  clientMutationId?: Maybe<Scalars['String']>;
};

export type NoopPayload = {
  __typename?: 'NoopPayload';
  clientMutationId?: Maybe<Scalars['String']>;
};

export type Query = {
  __typename?: 'Query';
  node?: Maybe<Node>;
  /** 「今日こと」を取得 */
  findWht: Array<Wht>;
};


export type QueryNodeArgs = {
  id: Scalars['ID'];
};


/** 今日こと */
export type Wht = Node & {
  __typename?: 'Wht';
  /** ID */
  id: Scalars['ID'];
  /** 記録日 */
  recordDate: Scalars['Date'];
  /** 画像パス */
  path: Scalars['String'];
};


/** 今日ことインプット */
export type WhtInput = {
  /** 記録日 */
  recordDate: Scalars['Date'];
  /** 画像 */
  image: Scalars['Upload'];
};
