/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface GormDeletedAt {
  time?: string;

  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface ModelsAnswer {
  answer_choices?: ModelsAnswerChoices[];
  answer_file?: string;
  answer_text?: string;
  candidate_id?: number;
  createdAt?: string;
  deletedAt?: GormDeletedAt;
  id?: number;
  point?: number;
  question_id: number;
  test_id?: number;
  updatedAt?: string;
}

export interface ModelsAnswerChoices {
  answer_id?: number;
  choices_id: number;
  createdAt?: string;
  deletedAt?: GormDeletedAt;
  id?: number;
  updatedAt?: string;
}

export interface ModelsCandidate {
  email?: string;
  id?: number;
  name?: string;
  test?: ModelsTest[];
}

export interface ModelsCandidateRequest {
  email?: string;
  name?: string;
  test?: ModelsTestResponse[];
}

export interface ModelsCandidateResponse {
  email?: string;
  id?: number;
  name?: string;
  test?: ModelsTest[];
}

export interface ModelsChoices {
  choice_text: string;
  is_answer: boolean;
}

export interface ModelsCreateQuestionInput {
  choices?: ModelsChoices[];
  difficulty: string;
  expected_time: number;
  file_read_me?: string;
  max_points: number;
  name: string;
  question_text: string;
  skill_id?: number;
  skill_name?: string;
  type?: string;
}

export interface ModelsCreateSkillInput {
  name: string;
}

export interface ModelsLoginInput {
  email: string;
  password: string;
}

export interface ModelsMyTests {
  expected_time?: number;
  number_Question?: number;
  number_candidate?: number;
  test_id?: number;
  test_name?: string;
}

export interface ModelsQuestion {
  answer?: ModelsAnswer[];
  choices?: ModelsChoices[];
  difficulty?: string;
  expected_time?: number;
  file_read_me?: string;
  id?: number;
  max_points?: number;
  name?: string;
  question_text?: string;
  skill?: ModelsSkill;
  skill_id?: number;
  test?: ModelsTest[];
  test_questions?: ModelsTestQuestion[];
  type?: string;
}

export interface ModelsSkill {
  id?: number;
  name?: string;
}

export interface ModelsStartTest {
  created_at?: string;
  current_question?: number;
  email?: string;
  name?: string;
  questions?: ModelsStartTestQuestions[];
  score?: number;
  test_status?: string;
  time_limit?: number;
}

export interface ModelsStartTestQuestions {
  expected_time?: number;
  name?: string;
  type?: string;
}

export interface ModelsTest {
  archived?: boolean;
  candidate?: ModelsCandidate[];
  description?: string;
  id?: number;
  name?: string;
  notify_emails?: string;
  passing_score?: number;
  questions?: ModelsQuestion[];
  show_score?: boolean;
  time_limit?: number;
  timing_policy?: string;
}

export interface ModelsTestCandidate {
  answer?: ModelsAnswer[];
  candidate_id: number;
  current_question?: number;
  score?: number;
  test_id: number;
  test_status?: string;
}

export interface ModelsTestQuestion {
  id?: number;
  question_id: number;
  test_id?: number;
}

export interface ModelsTestRequest {
  archived?: boolean;
  description?: string;
  name?: string;
  notify_emails?: string;
  passing_score?: number;
  show_score?: boolean;
  timing_policy?: string;
}

export interface ModelsTestResponse {
  archived?: boolean;
  description?: string;
  id?: number;
  name?: string;
  notify_emails?: string;
  passing_score?: number;
  show_score?: boolean;
  timing_policy?: string;
}

export interface ModelsTestsCandidatesResponse {
  candidate_email?: string;
  candidate_name?: string;
  score?: number;
  test_candidate_id?: string;
  test_name?: string;
  test_status?: string;
}

export interface ModelsUpdateCurrentQuestion {
  current_question?: number;
}

export interface ModelsUpdateTestStatus {
  test_status?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType | null) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "/api";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) => fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  private addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return `${value.map(this.addQueryParam).join("&")}`;
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) => (Array.isArray(query[key]) ? this.addArrayQueryParam(query, key) : this.addQueryParam(query, key)))
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title tekab-test
 * @version 1.0
 * @baseUrl /api
 * @contact
 *
 * this is an application web of interview assessment tests for interviewing out of the box .
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  answers = {
    /**
     * @description create new Answer by json
     *
     * @tags Answer
     * @name AnswersCreate
     * @summary add new Answer
     * @request POST:/answers/{idTestCandidate}
     */
    answersCreate: (idTestCandidate: string, Answer: ModelsAnswer, params: RequestParams = {}) =>
      this.request<ModelsAnswer, any>({
        path: `/answers/${idTestCandidate}`,
        method: "POST",
        body: Answer,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  candidate = {
    /**
     * @description create new Candidate by json
     *
     * @tags Candidate
     * @name CandidateCreate
     * @summary add new  Candidate
     * @request POST:/candidate
     */
    candidateCreate: (candidate: ModelsCandidateRequest[], params: RequestParams = {}) =>
      this.request<ModelsCandidateResponse, any>({
        path: `/candidate`,
        method: "POST",
        body: candidate,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  login = {
    /**
     * @description Login to the app
     *
     * @tags Login
     * @name LoginCreate
     * @summary Login to the app
     * @request POST:/login
     */
    loginCreate: (user: ModelsLoginInput, params: RequestParams = {}) =>
      this.request<string, any>({
        path: `/login`,
        method: "POST",
        body: user,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  myTests = {
    /**
     * @description create new Test by json
     *
     * @tags test
     * @name MyTestsCreate
     * @summary add new Test
     * @request POST:/my-tests
     */
    myTestsCreate: (Test: ModelsTestRequest, params: RequestParams = {}) =>
      this.request<ModelsTestResponse, any>({
        path: `/my-tests`,
        method: "POST",
        body: Test,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description create test_candidate by json and path
     *
     * @tags test_candidate
     * @name CandidatesTestIdCreate
     * @summary add new  test_candidate
     * @request POST:/my-tests/candidates/:test_id
     */
    candidatesTestIdCreate: (testId: number, test_candidate: ModelsTestCandidate, params: RequestParams = {}) =>
      this.request<ModelsTestCandidate, any>({
        path: `/my-tests/candidates/${testId}`,
        method: "POST",
        body: test_candidate,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description get my-tests
     *
     * @tags test
     * @name GetMyTests
     * @summary update Test
     * @request GET:/my-tests/getTest
     */
    getMyTests: (params: RequestParams = {}) =>
      this.request<ModelsMyTests, any>({
        path: `/my-tests/getTest`,
        method: "GET",
        format: "json",
        ...params,
      }),

    /**
     * @description delete a question from test by json
     *
     * @tags question_test
     * @name QuestionsDelete
     * @summary delete a question from test
     * @request DELETE:/my-tests/questions/{id}
     */
    questionsDelete: (id: string, params: RequestParams = {}) =>
      this.request<string, any>({
        path: `/my-tests/questions/${id}`,
        method: "DELETE",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description add a question to test by json
     *
     * @tags question_test
     * @name QuestionsCreate
     * @summary add a question to test
     * @request POST:/my-tests/{id}/questions
     */
    questionsCreate: (id: string, test_question: ModelsTestQuestion, params: RequestParams = {}) =>
      this.request<ModelsTestQuestion, any>({
        path: `/my-tests/${id}/questions`,
        method: "POST",
        body: test_question,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description update Test by json and path
     *
     * @tags test
     * @name UpdateTest
     * @summary update Test
     * @request POST:/my-tests/{test_id}
     */
    updateTest: (testId: number, Test: ModelsTestRequest, params: RequestParams = {}) =>
      this.request<ModelsTestResponse, any>({
        path: `/my-tests/${testId}`,
        method: "POST",
        body: Test,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  questions = {
    /**
     * @description find a question by type or difficulty
     *
     * @tags question
     * @name QuestionsList
     * @summary find a question
     * @request GET:/questions
     */
    questionsList: (query?: { type?: string; difficulty?: string }, params: RequestParams = {}) =>
      this.request<ModelsQuestion[], any>({
        path: `/questions`,
        method: "GET",
        query: query,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description create new question by json
     *
     * @tags question
     * @name EditCreate
     * @summary add new  question
     * @request POST:/questions/edit
     */
    editCreate: (question: ModelsCreateQuestionInput, params: RequestParams = {}) =>
      this.request<ModelsQuestion, any>({
        path: `/questions/edit`,
        method: "POST",
        body: question,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  quiz = {
    /**
     * @description Update current question
     *
     * @name CurrentQuestionUpdate
     * @summary Update current question
     * @request PUT:/quiz/currentQuestion/{idTestCandidate}
     */
    currentQuestionUpdate: (
      idTestCandidate: string,
      currentQuestion: ModelsUpdateCurrentQuestion,
      params: RequestParams = {},
    ) =>
      this.request<ModelsUpdateCurrentQuestion, any>({
        path: `/quiz/currentQuestion/${idTestCandidate}`,
        method: "PUT",
        body: currentQuestion,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description Update test status
     *
     * @name StatusUpdate
     * @summary Update test status
     * @request PUT:/quiz/status/{idTestCandidate}
     */
    statusUpdate: (idTestCandidate: string, testStatus: ModelsUpdateTestStatus, params: RequestParams = {}) =>
      this.request<ModelsUpdateTestStatus, any>({
        path: `/quiz/status/${idTestCandidate}`,
        method: "PUT",
        body: testStatus,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description get questions of a test
     *
     * @tags test
     * @name QuizDetail
     * @summary get tests candidates
     * @request GET:/quiz/{idTestCandidate}
     * @secure
     */
    quizDetail: (idTestCandidate: string, params: RequestParams = {}) =>
      this.request<ModelsQuestion[], any>({
        path: `/quiz/${idTestCandidate}`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  score = {
    /**
     * @description calculate score by path  and update a status test
     *
     * @tags test_candidate
     * @name ScoreCreate
     * @summary calculate a test score
     * @request POST:/score/{idTestCandidate}
     */
    scoreCreate: (idTestCandidate: string, params: RequestParams = {}) =>
      this.request<ModelsTestCandidate, any>({
        path: `/score/${idTestCandidate}`,
        method: "POST",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  skill = {
    /**
     * @description create new skill by json
     *
     * @tags skill
     * @name SkillCreate
     * @summary add new  skill
     * @request POST:/skill
     */
    skillCreate: (skill: ModelsCreateSkillInput, params: RequestParams = {}) =>
      this.request<ModelsSkill, any>({
        path: `/skill`,
        method: "POST",
        body: skill,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  skills = {
    /**
     * @description get all skills
     *
     * @tags skill
     * @name SkillsList
     * @summary get skills
     * @request GET:/skills
     * @secure
     */
    skillsList: (params: RequestParams = {}) =>
      this.request<ModelsSkill[], any>({
        path: `/skills`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  startTest = {
    /**
     * @description get test information
     *
     * @tags test_candidate
     * @name StartTest
     * @summary get test information
     * @request GET:/startTest/{idTestCandidate}
     * @secure
     */
    startTest: (idTestCandidate: string, params: RequestParams = {}) =>
      this.request<ModelsStartTest[], any>({
        path: `/startTest/${idTestCandidate}`,
        method: "GET",
        secure: true,
        format: "json",
        ...params,
      }),
  };
  tests = {
    /**
     * @description get tests by skill
     *
     * @tags test
     * @name TestsList
     * @summary get tests
     * @request GET:/tests
     * @secure
     */
    testsList: (query?: { type?: number }, params: RequestParams = {}) =>
      this.request<ModelsTestResponse[], any>({
        path: `/tests`,
        method: "GET",
        query: query,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  testscandidates = {
    /**
     * @description get candidates and their tests
     *
     * @tags test_candidate
     * @name TestscandidatesList
     * @summary get candidates and their tests
     * @request GET:/testscandidates
     * @secure
     */
    testscandidatesList: (params: RequestParams = {}) =>
      this.request<ModelsTestsCandidatesResponse[], any>({
        path: `/testscandidates`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
