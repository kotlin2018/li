export default {
  title: "Li Demo",
  copyright: "Powered by ❤️璃❤️",
  home: "WelcomePage",
  navitems: [
    {
      type: "void",
      name: "themeswitch",
      "x-component": "ThemeSwitch",
    },
    {
      type: "void",
      name: "currentuser",
      "x-component": "DropdownMenu",
      "x-component-props": {
        triggerProps: {
          popupStyle: {
            marginRight: "28px",
          },
        },
      },
      items: {
        type: "void",
        properties: {
          editprofile: {
            type: "void",
            "x-component": "DropdownMenu.Item",
            properties: {
              action: {
                title: "用户设置",
                type: "void",
                "x-component": "Action",
                "x-component-props": {
                  type: "menu",
                },
                properties: {
                  drawer1: {
                    "x-component": "Action.Drawer",
                    type: "void",
                    title: "Drawer Title",
                    properties: {
                      hello1: {
                        "x-content": "Hello",
                        title: "T1",
                      },
                      input: {
                        type: "string",
                        title: "Input",
                        "x-component": "Input",
                      },
                      footer1: {
                        "x-component": "Action.Drawer.Footer",
                        type: "void",
                        properties: {
                          close1: {
                            title: "Close",
                            "x-component": "Action",
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
          },
          changepwd: {
            type: "void",
            title: "修改密码",
            "x-component": "DropdownMenu.Item",
          },
          switchlang: {
            type: "void",
            title: "切换语言",
            "x-component": "DropdownMenu.SubMenu",
            properties: {
              switchlang2zh: {
                type: "void",
                title: "中文",
                "x-component": "DropdownMenu.Item",
              },
              switchlang2en: {
                type: "void",
                title: "Englist",
                "x-component": "DropdownMenu.Item",
              },
            },
          },
          divider: {
            type: "void",
            "x-component": "Divider",
            "x-component-props": {
              style: {
                margin: "4px 0",
              },
            },
          },
          signout: {
            type: "void",
            title: "退出登录",
            "x-component": "DropdownMenu.Item",
          },
        },
      },
      properties: {
        curUser: {
          type: "void",
          "x-component": "Avatar",
          "x-component-props": {
            size: 32,
            style: {
              cursor: "pointer",
            },
            src: "https://lf1-xgcdn-tos.pstatp.com/obj/vcloud/vadmin/start.8e0e4855ee346a46ccff8ff3e24db27b.png",
          },
        },
      },
    },
  ],
  menus: [
    {
      key: "WelcomePage",
      title: "欢迎",
    },
    {
      key: "AdminPage",
      title: "管理页",
      children: [
        {
          key: "AdminPageSub1",
          title: "一级页面",
        },
        {
          key: "AdminPageSub2",
          title: "二级页面",
        },
        {
          key: "AdminPageSub3",
          title: "三级页面",
        },
      ],
    },
    {
      title: "列表页",
      key: "ListPage",
      children: [
        {
          key: "AdminPageSub1Sub",
          title: "一级列表页面",
          children: [
            {
              key: "AdminPageSub1Sub1",
              title: "一一级列表页面",
            },
            {
              key: "AdminPageSub1Sub2",
              title: "一二级列表页面",
            },
            {
              key: "sAdminPageSub1Sub3",
              title: "一三级列表页面",
            },
          ],
        },
        {
          key: "AdminPageSub2",
          title: "二级列表页面",
        },
        {
          key: "AdminPageSub3",
          title: "三级列表页面",
        },
      ],
    },
    {
      target: "https://github.com/BeanWei",
      title: "Github",
    },
  ],
  binding: {
    signpage: {
      type: "object",
      "x-component": "Form",
      properties: {
        email: {
          type: "string",
          required: true,
          "x-component": "Input",
          "x-validator": "email",
          "x-decorator": "FormItem",
          "x-component-props": { placeholder: "邮箱", style: {} },
        },
        password: {
          type: "string",
          "x-component": "Password",
          required: true,
          "x-decorator": "FormItem",
          "x-component-props": { placeholder: "密码", style: {} },
        },
        submit: {
          type: "void",
          "x-content": "登录",
          "x-component": "Submit",
          "x-decorator": "FormItem",
          "x-operation": "userSignIn",
          "x-component-props": {
            block: true,
            type: "primary",
            style: { width: "100%" },
            onSubmitSuccessTo: "/admin/WelcomePage",
          },
        },
      },
    },
  },
};
