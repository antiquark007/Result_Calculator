import LightModeIcon from "@mui/icons-material/LightMode";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import {
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  useTheme,
  Box,
  Button,
} from "@mui/material";
import { useCallback, useContext, useEffect, useState } from "react";
import { StyleThemeContext } from "../../context/StyleThemeContext";
import { useNavigate } from "react-router-dom";

type Props = {};

const Appbar = (props: Props) => {
  const theme = useTheme();
  const navigate = useNavigate();

  const { toggleColorMode } = useContext(StyleThemeContext);
  const onClick = () => {
    navigate("/");
  };


  return (
    <AppBar
      position="fixed"
      sx={{
        width: "100%",
        zIndex: "1400",
        bgcolor: theme.palette.background.paper,
        color: theme.palette.text.primary,
      }}
    >
      <Toolbar>
        <Box
          sx={{ display: "flex", alignItems: "center", cursor: "pointer" }}
          onClick={onClick}
        >
          <Typography
            variant="h6"
            noWrap
            component="div"
            sx={{
              ml: 2,
            }}
          >
            Result Calculator
          </Typography>
        </Box>
        <Box sx={{ ml: "auto" }}>
          <IconButton onClick={toggleColorMode} color="inherit">
            {theme.palette.mode === "dark" ? (
              <DarkModeIcon />
            ) : (
              <LightModeIcon />
            )}
          </IconButton>
        </Box>
      </Toolbar>
    </AppBar>
  );
};

export default Appbar;
