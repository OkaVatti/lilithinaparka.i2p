export * from "./api";

export interface ThemeOption {
  name: string;
  value: string;
  preview: {
    bg: string;
    fg: string;
    primary: string;
  };
}
