import { App, Plugin, PluginSettingTab, Setting } from 'obsidian';

interface OBzhoushouSettings {
  cloudServerUrl: string;
  deviceCode: string;
  isPaired: boolean;
}

const DEFAULT_SETTINGS: OBzhoushouSettings = {
  cloudServerUrl: 'http://localhost:8080',
  deviceCode: '',
  isPaired: false,
};

export default class OBzhoushouPlugin extends Plugin {
  settings: OBzhoushouSettings;

  async onload() {
    await this.loadSettings();
    console.log('OBzhushou plugin loaded');
    
    // 添加状态栏按钮
    this.addStatusBarItem().setText('OBzhushou: Ready');

    // 添加命令
    this.addCommand({
      id: 'pair-device',
      name: 'Pair Device',
      callback: () => {
        console.log('Pair device clicked');
      },
    });

    this.addSettingTab(new OBzhoushouSettingTab(this.app, this));
  }

  onunload() {
    console.log('OBzhushou plugin unloaded');
  }

  async loadSettings() {
    this.settings = Object.assign({}, DEFAULT_SETTINGS, await this.loadData());
  }

  async saveSettings() {
    await this.saveData(this.settings);
  }
}

class OBzhoushouSettingTab extends PluginSettingTab {
  plugin: OBzhoushouPlugin;

  constructor(app: App, plugin: OBzhoushouPlugin) {
    super(app, plugin);
    this.plugin = plugin;
  }

  display(): void {
    const { containerEl } = this;

    containerEl.empty();

    new Setting(containerEl)
      .setName('Cloud Server URL')
      .setDesc('URL of the OBzhushou cloud service')
      .addText(text =>
        text
          .setPlaceholder('http://localhost:8080')
          .setValue(this.plugin.settings.cloudServerUrl)
          .onChange(async (value) => {
            this.plugin.settings.cloudServerUrl = value;
            await this.plugin.saveSettings();
          })
      );
  }
}
