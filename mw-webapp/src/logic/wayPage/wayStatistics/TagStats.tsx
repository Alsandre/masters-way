import {Fragment} from "react";
import {JobTagStat} from "src/logic/wayPage/wayStatistics/JobTagStat";
import {StatisticLine} from "src/logic/wayPage/wayStatistics/StatisticLine";

/**
 * TagStats props
 */
interface TagStatsProps {

  /**
   * Tags stats
   */
  stats: JobTagStat[];
}

/**
 * Render stats related to job tags
 */
export const TagStats = (props: TagStatsProps) => {
  return props.stats.map((tagStat) => (
    <Fragment key={tagStat.name}>
      <StatisticLine
        description={`${tagStat.name} jobs finished:`}
        value={`${tagStat.totalAmount} / ${tagStat.totalAmountPercentage} %`}
      />
      <StatisticLine
        key={tagStat.name}
        description={`${tagStat.name} time with:`}
        value={`${tagStat.totalTime}  / ${tagStat.totalTimePercentage} %`}
      />
    </Fragment>
  ));
};