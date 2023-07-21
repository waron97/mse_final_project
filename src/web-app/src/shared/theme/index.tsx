import { ThemeProvider as StyledThemeProvider } from "styled-components";

const theme = {
  colors: {
    primary: "#0070f3",
  },
  spacing: {},
};

export type Theme = { theme: typeof theme };

type Props = {
  children: React.ReactNode;
};

export default function ThemeProvider(props: Props) {
  return (
    <StyledThemeProvider theme={theme}>{props.children}</StyledThemeProvider>
  );
}
