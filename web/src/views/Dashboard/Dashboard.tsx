import React, { MouseEventHandler } from "react";

import { Typography, Grid, AppBar, Toolbar, IconButton, Drawer, Divider, Box } from "@material-ui/core";
import { createStyles, makeStyles, useTheme, Theme } from "@material-ui/core/styles";
import {
    DashboardOutlined,
    MenuOutlined,
    PeopleOutlined,
    SecurityOutlined,
    SettingsOutlined,
} from "@material-ui/icons";
import clsx from "clsx";

import AutheliaTitleIcon from "@assets/images/authelia-title.png";

const drawerWidth = 150;

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
            margin: theme.spacing(7),
            width: drawerWidth,
        },
        icon: {
            width: "inherit",
        },
        drawerElement: {
            color: "rgba(0,0,0,0.6)",
            display: "flex",
            alignItems: "center",
            marginLeft: theme.spacing(2),
            marginTop: theme.spacing(1),
            marginBottom: theme.spacing(1),
            cursor: "pointer",
        },
        drawerElementActive: {
            borderRight: `4px solid ${theme.palette.primary.main}`,
            color: `${theme.palette.primary.main}`,
        },
        drawerElementIcon: {
            marginRight: theme.spacing(1),
            width: "1.5em",
            height: "1.5em",
        },
    }),
);

enum Page {
    Dashboard = "Dashboard",
    Users = "Users",
    ACL = "ACL",
    Config = "Config",
}

const pageIcons = {
    [Page.Dashboard]: DashboardOutlined,
    [Page.Users]: PeopleOutlined,
    [Page.ACL]: SecurityOutlined,
    [Page.Config]: SettingsOutlined,
};

export interface DrawerElementProps {
    page: Page;
    active: boolean;
    onClick: MouseEventHandler;
}

const DrawerElement = function (props: DrawerElementProps) {
    const theme = useTheme();
    const classes = useStyles(theme);

    const Icon = pageIcons[props.page];

    return (
        <div
            className={clsx(classes.drawerElement, { [classes.drawerElementActive]: props.active })}
            onClick={props.onClick}
        >
            <Icon className={clsx(classes.drawerElementIcon)} />
            <Typography variant="h4">{Page[props.page]}</Typography>
        </div>
    );
};

const DashboardPage = () => {
    return <Box />;
};

const UsersPage = () => {
    return <Box />;
};

const ACLPage = () => {
    return <Box />;
};

const ConfigPage = () => {
    return <Box />;
};

const pages = {
    [Page.Dashboard]: DashboardPage,
    [Page.Users]: UsersPage,
    [Page.ACL]: ACLPage,
    [Page.Config]: ConfigPage,
};

const Dashboard = function () {
    const theme = useTheme();
    const classes = useStyles(theme);
    const [open, setOpen] = React.useState(false);
    const [page, setPage] = React.useState(Page.Dashboard);

    const handleDrawerOpen = () => {
        setOpen(true);
    };

    const handleDrawerClose = () => {
        setOpen(false);
    };

    const Content = pages[page];

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
                <MenuOutlined />
            </IconButton>
            <Drawer variant="persistent" anchor="left" open={open}>
                <div className={clsx(classes.iconContainer)}>
                    <img src={AutheliaTitleIcon} className={clsx(classes.icon)} />
                </div>

                <DrawerElement
                    page={Page.Dashboard}
                    active={page === Page.Dashboard}
                    onClick={() => {
                        setPage(Page.Dashboard);
                    }}
                />
                <DrawerElement
                    page={Page.Users}
                    active={page === Page.Users}
                    onClick={() => {
                        setPage(Page.Users);
                    }}
                />
                <DrawerElement
                    page={Page.ACL}
                    active={page === Page.ACL}
                    onClick={() => {
                        setPage(Page.ACL);
                    }}
                />
                <DrawerElement
                    page={Page.Config}
                    active={page === Page.Config}
                    onClick={() => {
                        setPage(Page.Config);
                    }}
                />
            </Drawer>
            <Content />
        </>
    );
};

export default Dashboard;
