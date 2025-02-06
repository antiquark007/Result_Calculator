import { Palette } from "@mui/material";

export const getDesignTokens = (mode: Palette["mode"]) => ({
  palette: {
    mode,
    ...(mode === "light"
      ? {
          // palette values for light mode
          primary: { main: "#75E6DA" },
          background: { default: "#189AB4", paper: "#05445E" },
        }
      : {
          // palette values for dark mode
          primary: { main: "#68BBE3" },
          background: { default: "#0C2D48", paper: "#0C2D48" },
        }),
  },
});
