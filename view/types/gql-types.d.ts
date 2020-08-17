export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  WhtID: any;
  Date: any;
  Upload: any;
  TextContentID: any;
  MovieContentID: any;
  ImageContentID: any;
  VoiceContentID: any;
};

/** 検索方法 */
export enum Compare {
  /** 完全一致 */
  Equal = 'Equal',
  /** 曖昧検索 */
  Like = 'Like'
}


/** 画像コンテンツ */
export type ImageContent = {
  __typename?: 'ImageContent';
  id: Scalars['ID'];
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** 画像パス */
  path: Scalars['String'];
};


/** 画像コンテンツインプット */
export type ImageContentInput = {
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** 画像 */
  image: Scalars['Upload'];
};

/** 動画コンテンツ */
export type MovieContent = {
  __typename?: 'MovieContent';
  id: Scalars['ID'];
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** 動画パス */
  path: Scalars['String'];
};


/** 動画コンテンツインプット */
export type MovieContentInput = {
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** 動画 */
  movie: Scalars['Upload'];
};

export type Mutation = {
  __typename?: 'Mutation';
  noop?: Maybe<NoopPayload>;
  /** 「今日こと」を登録 */
  createWht?: Maybe<MutationResponse>;
  /** テキストコンテンツを登録 */
  createTextContents?: Maybe<MutationResponse>;
  /** 画像コンテンツを登録 */
  createImageContents?: Maybe<MutationResponse>;
  /** 音声コンテンツを登録 */
  createVoiceContents?: Maybe<MutationResponse>;
  /** 動画コンテンツを登録 */
  createMovieContents?: Maybe<MutationResponse>;
};


export type MutationNoopArgs = {
  input?: Maybe<NoopInput>;
};


export type MutationCreateWhtArgs = {
  wht: WhtInput;
};


export type MutationCreateTextContentsArgs = {
  recordDate: Scalars['Date'];
  inputs: Array<TextContentInput>;
};


export type MutationCreateImageContentsArgs = {
  recordDate: Scalars['Date'];
  inputs: Array<ImageContentInput>;
};


export type MutationCreateVoiceContentsArgs = {
  recordDate: Scalars['Date'];
  inputs: Array<VoiceContentInput>;
};


export type MutationCreateMovieContentsArgs = {
  recordDate: Scalars['Date'];
  inputs: Array<MovieContentInput>;
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


export type QueryFindWhtArgs = {
  condition?: Maybe<WhtConditionInput>;
};

export enum Role {
  Admin = 'ADMIN',
  Manager = 'MANAGER',
  Editor = 'EDITOR',
  Viewer = 'VIEWER',
  Guest = 'GUEST'
}

/** テキストコンテンツ */
export type TextContent = {
  __typename?: 'TextContent';
  id: Scalars['ID'];
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** テキスト */
  text: Scalars['String'];
};


/** テキストコンテンツインプット */
export type TextContentInput = {
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** テキスト */
  text: Scalars['String'];
};


/** 音声コンテンツ */
export type VoiceContent = {
  __typename?: 'VoiceContent';
  id: Scalars['ID'];
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** 音声パス */
  path: Scalars['String'];
};


/** 音声コンテンツインプット */
export type VoiceContentInput = {
  /** コンテンツ名 */
  name?: Maybe<Scalars['String']>;
  /** 音声 */
  voice: Scalars['Upload'];
};

/** 今日こと */
export type Wht = Node & {
  __typename?: 'Wht';
  /** ID */
  id: Scalars['ID'];
  /** 記録日 */
  recordDate: Scalars['Date'];
  /** タイトル */
  title?: Maybe<Scalars['String']>;
  /** テキストコンテンツ */
  textContents?: Maybe<Array<TextContent>>;
  /** 画像コンテンツ */
  imageContents?: Maybe<Array<ImageContent>>;
  /** 画像コンテンツ */
  voiceContents?: Maybe<Array<VoiceContent>>;
  /** 動画コンテンツ */
  movieContents?: Maybe<Array<MovieContent>>;
};

/** 「今日こと」検索条件 */
export type WhtConditionInput = {
  /** ID */
  id?: Maybe<Scalars['WhtID']>;
  /** 記録日 */
  recordDate?: Maybe<Scalars['Date']>;
  /** タイトル */
  title?: Maybe<Scalars['String']>;
  /** 検索方法 */
  compare?: Maybe<Compare>;
};


/** 今日ことインプット */
export type WhtInput = {
  /** 記録日 */
  recordDate: Scalars['Date'];
  /** タイトル */
  title?: Maybe<Scalars['String']>;
};
