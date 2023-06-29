declare module 'express-serve-static-core' {
  interface Request {
    pagination: {
      skip: number;
      limit: number;
      page: number;
      pageSize: number;
    };
    user: {
      email: string;
      displayName: string;
      id: string;
    };
    registerCachedContent?: (content) => void;
  }
}

export {};
