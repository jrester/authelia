import React from "react";

import { Typography, Grid, AppBar, Toolbar, IconButton, Drawer, Divider } from "@material-ui/core";
import { createStyles, makeStyles, useTheme, Theme } from "@material-ui/core/styles";
import DashboardIcon from "@material-ui/icons/Dashboard";
import MenuIcon from "@material-ui/icons/Menu";
import clsx from "clsx";

import AutheliaTitleIcon from "@assets/images/authelia-title.png";

const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
        root: {
            display: "flex",
        },
        appBar: {
            zIndex: theme.zIndex.drawer + 1,
            transition: theme.transitions.create(["width", "margin"], {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen,
            }),
        },
        appBarShift: {
            marginLeft: drawerWidth,
            width: `calc(100% - ${drawerWidth}px)`,
            transition: theme.transitions.create(["width", "margin"], {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
        },
        menuButton: {
            marginLeft: 12,
            marginTop: 12,
        },
        hide: {
            display: "none",
        },
        drawer: {
            width: drawerWidth,
            flexShrink: 0,
            whiteSpace: "nowrap",
        },
        drawerOpen: {
            width: drawerWidth,
            transition: theme.transitions.create("width", {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
        },
        drawerClose: {
            transition: theme.transitions.create("width", {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.leavingScreen,
            }),
            overflowX: "hidden",
            width: theme.spacing(7) + 1,
            [theme.breakpoints.up("sm")]: {
                width: theme.spacing(9) + 1,
            },
        },
        toolbar: {
            display: "flex",
            alignItems: "center",
            justifyContent: "flex-end",
            padding: theme.spacing(0, 1),
            // necessary for content to be below app bar
            ...theme.mixins.toolbar,
        },
        content: {
            flexGrow: 1,
            padding: theme.spacing(3),
        },
        iconContainer: {
            marginLeft: "auto",
            marginRight: "auto",
            marginTop: 10,
            width: drawerWidth,
        },
        icon: {
            width: "-webkit-fill-available",
        },
    }),
);

enum Page {
    Dashboard = "Dashboard",
    Users = "Users",
    ACL = "ACL",
    Config = "Config",
}

export interface DrawerElementProps {
    page: Page;
    active: boolean;
}

const DrawerElement = function (props: DrawerElementProps) {
    const iconForElement = () => {
        switch (props.page) {
            case Page.Dashboard:
                return DashboardIcon;
            case Page.Users:
            case Page.ACL:
            case Page.Config:
        }
    };
    return (
        <div>
            <Typography variant="h2">{Page[props.page]}</Typography>
        </div>
    );
};

const Dashboard = function () {
    const theme = useTheme();
    const classes = useStyles();
    const [open, setOpen] = React.useState(false);
    const [content, setContent] = React.useState(Page.Dashboard);

    const handleDrawerOpen = () => {
        setOpen(true);
    };

    const handleDrawerClose = () => {
        setOpen(false);
    };

    return (
        <>
            <IconButton
                color="inherit"
                aria-label="open drawer"
                onClick={handleDrawerOpen}
                edge="start"
                className={clsx(classes.menuButton, {
                    [classes.hide]: open,
                })}
            >
                <MenuIcon />
            </IconButton>
            <Drawer variant="persistent" anchor="left" open={open}>
                <div className={clsx(classes.iconContainer)}>
                    <img src={AutheliaTitleIcon} className={clsx(classes.icon)} />
                </div>

                <DrawerElement page={Page.Dashboard} active={true} />
            </Drawer>
        </>
    );
};

export default Dashboard;
