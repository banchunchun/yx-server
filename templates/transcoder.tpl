<?xml version="1.0" encoding="UTF-8"?>
<TranscoderTask ID="{{.TaskId}}">
    <AVSettings>
        <VideoSettings Count="1">
            <VideoSetting idx="0">
                <ProcessMode>0</ProcessMode>
                <CodecType>{{.Codec}}</CodecType>
                <Encoding>Custom</Encoding>
                <Profile>High</Profile>
                <Level>Auto</Level>
                <Width>{{.Width}}</Width>
                <Height>{{.Height}}</Height>
                <FrameRate>2500</FrameRate>
                <FrameRateConversionMode>1</FrameRateConversionMode>
                <FrameRateSourceMode>0</FrameRateSourceMode>
                <RC>CBR</RC>
                <Quantizer>0</Quantizer>
                <BitRate>{{.BitRate}}</BitRate>
                <MaxBitRate>0</MaxBitRate>
                <VBVSize>175</VBVSize>
                <VBVDelay>700</VBVDelay>
                <GopSize>25</GopSize>
                <GopType>0</GopType>
                <Scenedetection>0</Scenedetection>
                <GopMode>1</GopMode>
                <BFrame>0</BFrame>
                <CABAC>1</CABAC>
                <Transform8x8>1</Transform8x8>
                <Intra8x8>1</Intra8x8>
                <LoopFilter>1</LoopFilter>
                <RefFrame>1</RefFrame>
                <Interlace>0</Interlace>
                <ThreadCount>5</ThreadCount>
                <LookHeadFrame>20</LookHeadFrame>
                <Policy>
                    <TwoPass>0</TwoPass>
                    <DeviceID>0</DeviceID>
                    <QualityLevel>2</QualityLevel>
                </Policy>
                <SmartStretch>
                    <DAR_X>-1</DAR_X>
                    <DAR_Y>-1</DAR_Y>
                    <Operate>1</Operate>
                    <FillColor>0</FillColor>
                </SmartStretch>
                <Deinterlace>1</Deinterlace>
                <Deblock>0</Deblock>
                <Delight>0</Delight>
                <Denoise>0</Denoise>
                <DenoiseMethod>0</DenoiseMethod>
                <Sharpen>0</Sharpen>
                <AntiAlias>0</AntiAlias>
                <AntiShaking>-1</AntiShaking>
                <EdgeAA>0</EdgeAA>
                <Bright>0</Bright>
                <Contrast>0</Contrast>
                <Hue>0</Hue>
                <Saturation>0</Saturation>
                <ReSizeAlg>3</ReSizeAlg>
                <DeinterlaceAlg>3</DeinterlaceAlg>
            </VideoSetting>
        </VideoSettings>
        <AudioSettings Count="1">
            <AudioSetting idx="0">
                <CodecType>{{.AudioAacType}}</CodecType>
                <Profile>LC</Profile>
                <PackageMode>1</PackageMode>
                <Channel>{{.AudioChannel}}</Channel>
                <BitRate>{{.AudioBitRate}}</BitRate>
                <SampleRate>{{.AudioSampleRate}}</SampleRate>
                <Denoise>0</Denoise>
                <BoostLevel>0</BoostLevel>
                <ChannelProcessing>None</ChannelProcessing>
                <VolumeProcessMode>0</VolumeProcessMode>
                <BalanceDB>-30</BalanceDB>
                <BalanceLevel>0</BalanceLevel>
            </AudioSetting>
        </AudioSettings>
    </AVSettings>
    <OutputGroups Count="1">
        <OutputGroup idx="0" label="">
            <EnableOutput>1</EnableOutput>
            <OutputType>FileArchive</OutputType>
            <OutputCount>1</OutputCount>
            <Container>{{.Container}}</Container>
            <TargetPath>{{.StreamOutput}}</TargetPath>
            <SegmentMode>0</SegmentMode>
            <SegmentRecordPara>
                <SRDataOffset>5</SRDataOffset>
                <EpgInfoFilePath></EpgInfoFilePath>
            </SegmentRecordPara>
            <MP4Setting>
                <FileHeadForemost>1</FileHeadForemost>
                <SplitAudio>0</SplitAudio>
                <SplitChannel>0</SplitChannel>
                <TargetPath>/</TargetPath>
                <TargetName>${r'${II}'}.wav</TargetName>
                <AutoStartEnd>0</AutoStartEnd>
            </MP4Setting>
            <Output idx="0" label="">
                <VideoSettingIdx>0</VideoSettingIdx>
                <AudioSettingIdx>0</AudioSettingIdx>
                <AudioInputIdx>-1</AudioInputIdx>
                <Subdirectory>0</Subdirectory>
                <HLSInfo>
                    <Type>0</Type>
                    <AudioGroupID></AudioGroupID>
                    <SubtitleGroupID></SubtitleGroupID>
                </HLSInfo>
            </Output>
        </OutputGroup>
    </OutputGroups>
    <Inputs Count="1">
        <Input idx="0">
        <ContentDetectType>3</ContentDetectType>
        <ContentDetectXml>{{.FeatureXmlInput}}</ContentDetectXml>
        <VideoDecoding>2</VideoDecoding>
        <Type>LocalFile</Type>
        <URI>{{.StreamInput}}</URI>
        <Program Id="-1">
            <VideoId>-1</VideoId>
            <AudioId>-1</AudioId>
            <SubtitleId>-3</SubtitleId>
        </Program>
        <Preprocessor>
            <Clips Count="0">
            </Clips>
            <AudioDelay>0</AudioDelay>
        </Preprocessor>
        </Input>
    </Inputs>
</TranscoderTask>
