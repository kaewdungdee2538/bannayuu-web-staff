module.exports = {
  apps: [
    {
      name: "bannayuu-web-admin",
      script: "./bannayuu-web-admin",
      env: {
        NODE_ENV: "production",
        DB_HOST: "local.uat.bannayuu.com",
        DB_NAME: "cit_bannayuu_db",
        DB_PORT: "5432",
        APP_PORT: ":4501",
        AUTHEN_TOKEN: "f56c3775-07b0-45e7-800f-304274533cb7",
        ROOT_IMAGE: "/home/ubuntu/bannayuu/",
        WEB_MANAGEMENT_RESET_USER: "https://cit.bannayuu.com/#/reset",
        DB_MAX_IDLE_TIME: "10",
        DB_MAX_CONECTIOS: "10",
      },
      env_development: {
        NODE_ENV: "development",
        DB_HOST: "local.uat.bannayuu.com",
        DB_NAME: "uat_cit_bannayuu_db",
        DB_PORT: "5432",
        APP_PORT: ":4501",
        AUTHEN_TOKEN: "f56c3775-07b0-45e7-800f-304274533cb7",
        ROOT_IMAGE: "/home/ubuntu/bannayuu/",
        WEB_MANAGEMENT_RESET_USER: "https://uat.bannayuu.com/#/reset",
        DB_MAX_IDLE_TIME: "10",
        DB_MAX_CONECTIOS: "10",
      },
    },
  ],
};
