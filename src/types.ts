// types.ts
export interface ParseResult {
  success: boolean;
  data: any[];
  error?: string;
}

export interface ParsedData {
  id: number;
  name: string;
  value: number;
}

export interface ParseOptions {
  delimiter: string;
  skipRows: number;
}

export type DataType = 'csv' | 'json' | 'xml';

export interface DataParserConfig {
  type: DataType;
  options: ParseOptions;
}